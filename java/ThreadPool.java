
import java.util.Collection;
import java.util.Queue;
import java.util.concurrent.ConcurrentLinkedQueue;

/**
 * Created by jiangKlijna on 2017/3/7.
 */
public class ThreadPool {

    private final int poolsize;
    private final WorkThread[] mWorkThreads;
    private final Queue mTaskQueue;

    public ThreadPool() {
        this(5);
    }

    public ThreadPool(int poolsize) {
        if (poolsize < 1) poolsize = 5;
        this.poolsize = poolsize;
        mWorkThreads = new WorkThread[poolsize];
        mTaskQueue = new ConcurrentLinkedQueue();
        for (int i = 0; i < poolsize; i++) {
            WorkThread wt = new WorkThread();
            wt.start();
            mWorkThreads[i] = wt;
        }
    }

    //执行任务,把任务加入队列，并notify，通知工作线程执行
    public void execute(Runnable task) {
        synchronized (mTaskQueue) {
            mTaskQueue.add(task);
            mTaskQueue.notify();
        }
    }

    public void execute(Runnable... tasks) {
        synchronized (mTaskQueue) {
            for (Runnable task : tasks)
                mTaskQueue.add(task);
            mTaskQueue.notify();
        }
    }

    public void execute(Collection<Runnable> tasks) {
        synchronized (mTaskQueue) {
            mTaskQueue.addAll(tasks);
            mTaskQueue.notify();
        }
    }

    //销毁线程池，清空任务队列，并停止所有工作线程
    public void destory() {
        mTaskQueue.clear();
        for (int i = 0; i < mWorkThreads.length; i++) {
            mWorkThreads[i].stopWorker();
            mWorkThreads[i] = null;
        }
    }

    //获得线程池的数量
    public int getPoolsize() {
        return poolsize;
    }

    //获得当前任务的数量
    public int getWaitTaskSize() {
        return mTaskQueue.size();
    }

    private class WorkThread extends Thread {
        private boolean isRunning = true;

        //若队列为空则继续等待，直到notify，取出队列头元素执行
        @Override
        public void run() {
            Runnable task;
            while (isRunning) {
                synchronized (mTaskQueue) {
                    while (isRunning && mTaskQueue.isEmpty()) {
                        try {
                            mTaskQueue.wait(2000);
                        } catch (InterruptedException e) {
                            e.printStackTrace();
                        }
                    }
                    task = (Runnable) mTaskQueue.poll();
                }
                if (task != null)
                    task.run();
            }
        }

        public void stopWorker() {
            isRunning = false;
        }
    }
}
