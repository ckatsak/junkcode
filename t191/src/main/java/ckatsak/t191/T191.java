package ckatsak.t191;

import org.apache.http.impl.client.CloseableHttpClient;
import org.apache.http.impl.client.HttpClients;
import org.apache.http.util.EntityUtils;
import org.apache.http.client.methods.HttpPost;
import org.apache.http.entity.mime.content.FileBody;
import org.apache.http.HttpEntity;
import org.apache.http.entity.ContentType;
import org.apache.http.entity.mime.MultipartEntityBuilder;
import org.apache.http.client.methods.CloseableHttpResponse;

import java.io.File;

public class T191 {

    private final String filePath;

    private final CloseableHttpClient c;

    T191(String filePath) {
        this.filePath = filePath;
        this.c = HttpClients.createDefault();
    }

    void doPOST(String targetURL) throws Exception {
        /*
        FileBody bin = new FileBody(new File(this.filePath));
        HttpEntity reqEntity = MultipartEntityBuilder.create()
            .addPart("file", bin)
            .build();
        */
        String[] filename = this.filePath.split("/", 0);
        HttpEntity reqEntity = MultipartEntityBuilder
            .create()
            //.addBinaryBody("file", new File(this.filePath), ContentType.APPLICATION_OCTET_STREAM, "JobGraphFile")
            .addBinaryBody("file",
                    new File(this.filePath),
                    ContentType.APPLICATION_OCTET_STREAM,
                    filename[filename.length - 1])
            .build();

        HttpPost p = new HttpPost(targetURL);
        p.setEntity(reqEntity);

        System.out.println("Executing request: " + p.getRequestLine());
        CloseableHttpResponse res = this.c.execute(p);
        try {
            System.out.println("Response: " + res.getStatusLine());
            HttpEntity resEntity = res.getEntity();
            if (resEntity != null) {
                System.out.println("Response's Content-length: " + resEntity.getContentLength());
            }
            EntityUtils.consume(resEntity);
        } finally {
            res.close();
        }
    }


    public static void main(String[] args) throws Exception {
        if (args.length != 2) {
            System.err.printf("\nArguments:\n\t<local file path>  <URL>\n\n");
            System.exit(1);
        }

        T191 t = new T191(args[0]);
        t.doPOST(args[1]);
    }

}
