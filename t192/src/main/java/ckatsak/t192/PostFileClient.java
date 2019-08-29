package ckatsak.t192;

import org.apache.http.impl.client.CloseableHttpClient;
import org.apache.http.impl.client.HttpClients;
import org.apache.http.util.EntityUtils;
import org.apache.http.client.methods.HttpPost;
import org.apache.http.HttpEntity;
import org.apache.http.entity.ContentType;
import org.apache.http.entity.mime.MultipartEntityBuilder;
import org.apache.http.client.methods.CloseableHttpResponse;

import java.io.File;

public class PostFileClient implements AutoCloseable {

    private final String filePath;

    private final CloseableHttpClient c;

    PostFileClient(final String filePath) {
        this.filePath = filePath;
        this.c = HttpClients.createDefault();
    }

    public void close() throws java.io.IOException {
        this.c.close();
    }

    void doPOST(final String targetURL) throws Exception {
        final HttpEntity reqEntity = MultipartEntityBuilder
            .create()
            .addBinaryBody("file",
                    new File(this.filePath),
                    ContentType.APPLICATION_OCTET_STREAM,
                    this.filePath.lastIndexOf('/') != -1
                        ? this.filePath.substring(this.filePath.lastIndexOf('/'))
                        : this.filePath)
            .build();

        final HttpPost p = new HttpPost(targetURL);
        p.setEntity(reqEntity);

        System.out.println("Executing request: " + p.getRequestLine());
        final CloseableHttpResponse res = this.c.execute(p);
        try {
            System.out.println("Response: " + res.getStatusLine());
            final HttpEntity resEntity = res.getEntity();
            if (resEntity != null) {
                System.out.println("Response's Content-length: " + resEntity.getContentLength());
            }
            EntityUtils.consume(resEntity);
        } finally {
            res.close();
        }
    }

}
