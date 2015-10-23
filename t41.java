import java.lang.System;
import java.lang.Thread;

public class t41 {

    public static void main(String[] args) {

        long t1, t2;

        t1 = System.currentTimeMillis();
        try {
            Thread.sleep(2500);
        } catch (InterruptedException ex) {
            Thread.currentThread().interrupt();
        }
        System.out.println("Beep: " + t1);

        t2 = System.currentTimeMillis();
        System.out.println("Beep: " + t2);

        System.out.println("Time inbetween: " + (t2 - t1) / 1000.0 + "sec");

    }
}
