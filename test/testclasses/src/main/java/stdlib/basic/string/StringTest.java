package stdlib.basic.string;

import static helper.MyAssert.*;

public class StringTest implements Runnable {

    public static void main(String[] args) {
        new StringTest().run();
    }

    @Override
    public void run() {
        String s1 = "abc1";
        String s2 = "abc1";
        assertSame(s1, s2);

        int x = 1;
        String s3 = new StringBuilder("abc").append(x).toString();
        assertNotSame(s1, s3);

        s3 = s3.intern();
        assertSame(s1, s3);
    }

}
