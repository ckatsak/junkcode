package ckatsak.t192;

public class T192 {

    public static void main(String[] args) throws Exception {
        if (args.length != 2) {
            System.err.printf("\nArguments:\n\t<local file path>  <URL>\n\n");
            System.exit(1);
        }

        try (final PostFileClient c = new PostFileClient(args[0])) {
            c.doPOST(args[1]);
        }
    }

}
