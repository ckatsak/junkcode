package gr.ntua.ece.cslab.ckatsak;

import java.util.logging.ConsoleHandler;
import java.util.logging.Logger;

public class A implements Runnable {

    private final Logger logger;// = Logger.getLogger(A.class.getCanonicalName());

    private final String name;

    public A(final String name) {
        this.name = name;
        this.logger = Logger.getLogger(A.class.getCanonicalName() + "--" + this.name);
        this.logger.addHandler(BQLogHandler.getInstance());

        //this.logger.addHandler(new ConsoleHandler());
        //this.logger.addHandler(new ConsoleHandler());
    }

    @Override
    public void run() {
        logger.info("entering " + this.name + ".run()!");

        for (int i = 0; i < 10; i++) {
            logger.info("[" + this.name + "] " + i);
            try {
                Thread.sleep(750);
            } catch (InterruptedException e) {
                logger.severe(e.getMessage());
                e.printStackTrace();
            }
        }

        logger.warning("exiting " + this.name + ".run()...");
    }
}
