import java.io.*;

public class Skata {
        private int[] yo;

        public Skata(int[] yo) {
                this.yo = yo;
        }

        public void print() {
                for (int i = 0; i < yo.length; i++) {
                        System.out.println(yo[i]);
                }
        }
}
