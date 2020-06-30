package gr.ntua.ece.cslab.ckatsak;

import java.util.concurrent.BlockingQueue;
import java.util.logging.ConsoleHandler;
import java.util.logging.Logger;

public class Consumer implements Runnable {

    private final Logger logger;
    private final String name;

    public Consumer(final String name) {
        this.name = name;
        this.logger = Logger.getLogger(Consumer.class.getCanonicalName() + "--" + this.name);
        this.logger.addHandler(BQLogHandler.getInstance());
        //this.logger.addHandler(new ConsoleHandler());
    }

    @Override
    public void run() {
        logger.info("entering " + this.name + ".run()!");

        final BlockingQueue<String> q = BQLogHandler.getQueue();

        for (int i = 5; i > 0; i--) {
            logger.info("[" + this.name + "] Polling the queue in " + i + "...");
            try {
                Thread.sleep(1000);
            } catch (InterruptedException e) {
                logger.severe(e.getMessage());
                e.printStackTrace();
            }
        }

        while (null != q.peek()) {
            final String logMsg = q.poll();
            System.out.println("log message retrieved: " + logMsg);
        }
        logger.warning("Queue is empty!");

        logger.info("leaving " + this.name + ".run()...");
    }
}
