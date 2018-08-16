
import java.util.ArrayList;
import java.util.List;
import java.util.concurrent.locks.Lock;
import java.util.concurrent.locks.ReentrantLock;

public class AlternateThread {

    private int state;
    private int count;
    private List<WorkThread> workers;
    private final Lock lock = new ReentrantLock();

    public AlternateThread(Runnable... runnables) {
        this.count = runnables.length;
        this.workers = new ArrayList<>(count);
        for (int i = 0; i < count; i++) {
            final Runnable runnable = runnables[i];
            workers.add(new WorkThread(i, () -> {
                runnable.run();
                return true;
            }));
        }
    }

    public AlternateThread(AlternateRunnable... runnables) {
        this.count = runnables.length;
        this.workers = new ArrayList<>(count);
        for (int i = 0; i < count; i++)
            workers.add(new WorkThread(i, runnables[i]));
    }

    public void remalloc(WorkThread garbage) {
        workers.remove(garbage);
        count--;
        for (int i = 0; i < count; i++)
            workers.get(i).index = i;
    }

    public void start() {
        workers.forEach(Thread::start);
    }

    private class WorkThread extends Thread {

        private int index;
        private final AlternateRunnable target;

        private WorkThread(int index, AlternateRunnable target) {
            this.index = index;
            this.target = target;
        }

        @Override
        public void run() {
            while (true)
                try {
                    lock.lock();
                    while (state % count == index) {
                        boolean isRun = target.run();
                        state++;
                        if (!isRun) {
                            remalloc(this);
                            return;
                        }
                    }
                } catch (Exception e) {
					remalloc(this);
                    throw new RuntimeException(e);
                } finally {
                    lock.unlock();
                }
        }
    }

    interface AlternateRunnable {
        // Run if run() return true
        // Stop if run() return false
        boolean run() throws Exception;
    }

    public static void main(String[] args) {
//        new AlternateThread(
//                () -> System.out.print("A"),
//                () -> System.out.print("B"),
//                () -> System.out.print("C"),
//                () -> System.out.print("D")
//        ).start();

//        long start = System.currentTimeMillis();
//        java.util.Random r = new java.util.Random();
//        new AlternateThread(
//                () -> {
//                    System.out.print("A");
//                    Thread.sleep(500);
//                    return true;
//                },
//                () -> {
//                    System.out.print("5");
//                    long end = System.currentTimeMillis();
//                    return end - start < 5000;
//                },
//                () -> {
//                    System.out.print("R");
//                    return r.nextBoolean();
//                },
//                () -> {
//                    System.out.print("Z");
//                    Thread.sleep(500);
//                    return true;
//                }
//        ).start();

        new AlternateThread(
                new AlternateRunnable() {
                    int i = 1;

                    @Override
                    public boolean run() throws Exception {
                        System.out.print(i++);
                        return i <= 26;
                    }
                },
                new AlternateRunnable() {
                    char c = 'a';

                    @Override
                    public boolean run() throws Exception {
                        System.out.print(c++);
                        return c <= 'z';
                    }
                }
        ).start();
    }

}
