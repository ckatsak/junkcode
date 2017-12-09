import java.io.*;

public class t149 {
        public static void main(String[] args) {
                int[] x = new int[10];
                for (int i = 0; i < 10; i++) {
                        x[i] = 1<<i;
                }

                System.out.println("Hello:");
                //for (int i = 0; i < 10; i++) {
                //        System.out.println(x[i]);
                //}
                Skata skata = new Skata(x);
                skata.print();
        }
}
