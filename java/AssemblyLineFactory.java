
import java.io.PrintStream;
import java.util.ArrayList;
import java.util.List;
import java.util.concurrent.*;

public class AssemblyLineFactory<T> {

    public static void main(String[] args) {
        AssemblyLineFactory<PrintStream> factory = new AssemblyLineFactory<>(
                out -> {
                    out.println(Thread.currentThread().getName());
                },
                out -> {
                    out.println(Thread.currentThread().getName());
                },
                out -> {
                    out.println(Thread.currentThread().getName());
                }
        );

        factory.start(System.out);
        factory.start(System.err);
    }

    private final List<Worker> workers;
    private final ExecutorService pool;

    public AssemblyLineFactory(Task<T>... tasks) {
        workers = new ArrayList<>(tasks.length);
        pool = Executors.newFixedThreadPool(tasks.length);
        for (int i = 0; i < tasks.length; i++) {
            Worker w = new Worker(i, tasks[i]);
            workers.add(w);
            pool.submit(w);
        }
    }

    public void start(T t) {
        workers.get(0).queue.add(t);
    }

    public void end(){
        pool.shutdown();
    }

    interface Task<T> {
        void run(T t);
    }

    private class Worker implements Runnable {
        private final int index;
        private final Task<T> task;
        private final BlockingQueue<T> queue;

        private Worker(int index, Task<T> task) {
            this.index = index;
            this.task = task;
            this.queue = new LinkedBlockingQueue<T>();
        }

        private Worker nextWorker() {
            if (index + 1 >= workers.size()) return null;
            return workers.get(index + 1);
        }

        @Override
        public void run() {
            while (true) {
                try {
                    T t = queue.take();
                    task.run(t);
                    Worker w = nextWorker();
                    if (w != null) w.queue.add(t);
                } catch (InterruptedException e) {
                }
            }
        }
    }
}
