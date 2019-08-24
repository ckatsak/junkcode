package ckatsak.helloworld;

import org.apache.commons.httpclient.*;
import org.apache.commons.httpclient.methods.*;
import org.apache.commons.httpclient.params.HttpMethodParams;

import java.io.IOException;

public class HelloWorld {

    private final String msg;

    private static final String url = "http://httpbin.org/get";
    private final HttpClient c;

    public HelloWorld(String msg) {
        this.msg = msg;
        this.c = new HttpClient();
    }

    public String getMsg() {
        return this.msg;
    }

    public static void main(String[] args) {
        HelloWorld hw = new HelloWorld("Hello, World!");

        System.out.println(hw.getMsg());

        GetMethod method = new GetMethod(url);
        //method.getParams().setParameter(HttpMethodParams.RETRY_HANDLER, new DefaultHttpMethodRetryHandler(3, false));

        try {
            int statusCode = hw.c.executeMethod(method);

            if (statusCode != HttpStatus.SC_OK) {
                System.err.println("Method failed: " + method.getStatusLine());
            }

            byte[] responseBody = method.getResponseBody();
            System.out.println("\nRESPONSE BODY:\n" + new String(responseBody) + "\n\n");
        } catch (HttpException e) {
            System.err.println("Fatal protocol violation: " + e.getMessage());
            e.printStackTrace();
        } catch (IOException e) {
            System.err.println("Fatal transport error: " + e.getMessage());
            e.printStackTrace();
        } finally {
            method.releaseConnection();
        }
    }

}
