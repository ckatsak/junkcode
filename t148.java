import java.io.*;

public class t148 {
        public static void main(String[] args) {
                int[] x = new int[10];
                for (int i = 0; i < 10; i++) {
                        x[i] = 1<<i;
                }

                System.out.println("Hello:");
                //for (int i = 0; i < 10; i++) {
                //        System.out.println(x[i]);
                //}
                skata skata = new skata(x);
                skata.print();
        }
}

class skata {
        private int[] yo;

        public skata(int[] yo) {
                this.yo = yo;
        }

        public void print() {
                for (int i = 0; i < yo.length; i++) {
                        System.out.println(yo[i]);
                }
        }
}
