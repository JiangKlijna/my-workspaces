
import java.util.Deque;
import java.util.concurrent.ConcurrentLinkedDeque;
import java.util.concurrent.atomic.AtomicInteger;

public class ProducerConsumer {

    public static class LaShiZhe implements Runnable {
        public final long millis;
        public final BlockingQueue<Shi> queue;

        public LaShiZhe(int second, BlockingQueue<Shi> queue) {
            this.millis = second * 1000;
            this.queue = queue;
        }

        public Shi la() {
            Shi shi = new Shi();
            System.out.println(this + " la " + shi + " bucket@" + queue.size());
            return shi;
        }

        @Override
        public void run() {
            while (true) {
                try {
                    queue.add(la());
                    Thread.sleep(millis);
                } catch (InterruptedException e) {
                    e.printStackTrace();
                }
            }
        }

        @Override
        public String toString() {
            return "LaShiZhe";
        }
    }

    public static class ChiShiZhe implements Runnable {
        public final long millis;
        public final BlockingQueue<Shi> queue;

        public ChiShiZhe(int second, BlockingQueue<Shi> queue) {
            this.millis = second * 1000;
            this.queue = queue;
        }

        public void chi(Shi shi) {
            System.out.println(this + " chi " + shi + " bucket@" + queue.size());
        }

        @Override
        public void run() {
            while (true) {
                try {
                    chi(queue.take());
                    Thread.sleep(millis);
                } catch (InterruptedException e) {
                    e.printStackTrace();
                }
            }
        }

        @Override
        public String toString() {
            return "ChiShiZhe";
        }
    }

    public static class Shi {
        private static final AtomicInteger i = new AtomicInteger();
        public final int tag = i.getAndIncrement();

        @Override
        public String toString() {
            return "Shi@" + tag;
        }
    }

    public static class BlockingQueue<E> {

        private final Deque<E> deque = new ConcurrentLinkedDeque<>();

        public int size() {
            return deque.size();
        }

        public void add(E e) {
            deque.add(e);
        }

        public E take() throws InterruptedException {
            while (true) {
                synchronized (deque) {
                    if (deque.isEmpty()) {
                        deque.wait(2000);
                    } else {
                        return deque.pollFirst();
                    }
                }
            }
        }

    }

    public static void main(String[] args) {
//        BlockingQueue<Shi> queue = new LinkedBlockingQueue<>();
        BlockingQueue<Shi> queue = new BlockingQueue<>();
        new Thread(new LaShiZhe(2, queue)).start();
        new Thread(new ChiShiZhe(3, queue)).start();
    }
}
