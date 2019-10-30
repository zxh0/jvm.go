package jvm.instructions;

import static helper.MyAssert.*;

public class LookupSwitchTest implements Runnable {

    public static void main(String[] args) {
        new LookupSwitchTest().run();
    }

    @Override
    public void run() {
        assertEquals("100", test(100));
        assertEquals("200", test(200));
        assertEquals("300", test(300));
        assertEquals("500", test(500));
        assertEquals("default", test(1));
    }

    private static String test(int x) {
        String str;
        switch (x) {
            case 100: str = "100"; break;
            case 200: str = "200"; break;
            case 300: str = "300"; break;
            case 500: str = "500"; break;
            default: str = "default";
        }
        return str;
    }

}
