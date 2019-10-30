package jvm.instructions;

import static helper.MyAssert.*;

public class TableSwitchTest implements Runnable {

    public static void main(String[] args) {
        new TableSwitchTest().run();
    }

    @Override
    public void run() {
        assertEquals("2", test(2));
        assertEquals("3", test(3));
        assertEquals("4", test(4));
        assertEquals("5", test(5));
        assertEquals("default", test(1));
    }

    private static String test(int x) {
        String str;
        switch (x) {
            case 2: str = "2"; break;
            case 3: str = "3"; break;
            case 4: str = "4"; break;
            case 5: str = "5"; break;
            default: str = "default";
        }
        return str;
    }

}
