package gr.ntua.ece.cslab.ckatsak;

import java.util.concurrent.BlockingQueue;
import java.util.concurrent.LinkedBlockingQueue;
import java.util.logging.ErrorManager;
import java.util.logging.Handler;
import java.util.logging.LogRecord;
import java.util.logging.Logger;

/**
 * TODO: Documentation
 */
public class BQLogHandler extends Handler {

    private static BQLogHandler handler;
    private static BlockingQueue<String> queue;

    private BQLogHandler() {
        super();
        queue = new LinkedBlockingQueue<>();
    }

    public static synchronized BQLogHandler getInstance() {
        if (null == handler) {
            handler = new BQLogHandler();
        }
        return handler;
    }

    public static synchronized BlockingQueue<String> getQueue() {
        getInstance();
        return queue;
    }


    //----------------------------------------------------------------------------------------------------------------


    @Override
    public void publish(final LogRecord logRecord) {
        if (!isLoggable(logRecord)) {
            return;
        }

        //queue.add(logRecord.getMessage());
        if (!queue.offer(logRecord.getMessage())) {
            getErrorManager().error("Queue is full", null, ErrorManager.WRITE_FAILURE);
        }
    }

    @Override
    public void flush() {
        // FIXME?
    }

    @Override
    public void close() throws SecurityException {
        // FIXME?
    }


    //----------------------------------------------------------------------------------------------------------------


    public static void main(final String[] args) {
        Logger.getLogger("").addHandler(BQLogHandler.getInstance());

        final A a0 = new A("a0");
        final A a1 = new A("a1");
        final A a2 = new A("a2");

        final Consumer c0 = new Consumer("consumer0");
        final Consumer c1 = new Consumer("consumer1");

        final Thread ta0 = new Thread(a0);
        final Thread ta1 = new Thread(a1);
        final Thread ta2 = new Thread(a2);
        final Thread tc0 = new Thread(c0);
        final Thread tc1 = new Thread(c1);

        tc0.start();
        ta0.start();
        ta1.start();
        ta2.start();

        try {
            tc0.join();
        } catch (InterruptedException e) {
            e.printStackTrace();
        }
        try {
            ta0.join();
            ta1.join();
            ta2.join();
        } catch (InterruptedException e) {
            e.printStackTrace();
        }
        tc1.start();
        try {
            tc1.join();
        } catch (InterruptedException e) {
            e.printStackTrace();
        }
    }

}
