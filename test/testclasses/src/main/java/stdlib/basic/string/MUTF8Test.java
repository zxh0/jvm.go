package stdlib.basic.string;

import static helper.MyAssert.*;

// http://www.oracle.com/technetwork/articles/javase/supplementary-142654.html
public class MUTF8Test implements Runnable {

    public static void main(String[] args) {
        new MUTF8Test().run();
    }

    @Override
    public void run() {
        assertEquals(1, "\u0000".length()); // NULL
        assertEquals(1, "A".length()); // U+0041
        //assertEquals(1, "ß".length()); // U+00DF
        assertEquals(1, "東".length()); // U+6771
        assertEquals(2, "\uD801\uDC00".length()); // U+10400
    }

}
