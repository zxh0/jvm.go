package main;

import jvm.instructions.LookupSwitchTest;
import jvm.instructions.TableSwitchTest;

public class AllTests {

    public static void main(String[] args) {
        runTest(new TableSwitchTest());
        runTest(new LookupSwitchTest());
    }

    private static void runTest(Runnable test) {
        test.run();
        System.out.println("[ok] " + test.getClass().getName());
    }

}
