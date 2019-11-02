package stdlib.basic;

import static helper.MyAssert.*;

public class StringTest implements Runnable {

    public static void main(String[] args) {
        new StringTest().run();
    }

    @Override
    public void run() {
        testIntern();
        testMUTF8();
    }

    private static void testIntern() {
        String s1 = "abc1";
        String s2 = "abc1";
        assertSame(s1, s2);

        int x = 1;
        String s3 = new StringBuilder("abc").append(x).toString();
        assertNotSame(s1, s3);

        s3 = s3.intern();
        assertSame(s1, s3);
    }

    private static void testMUTF8() {
        assertEquals(1, "\u0000".length()); // NULL
        assertEquals(1, "A".length()); // U+0041
        //assertEquals(1, "ß".length()); // U+00DF
        assertEquals(1, "東".length()); // U+6771
        assertEquals(2, "\uD801\uDC00".length()); // U+10400
    }

}
