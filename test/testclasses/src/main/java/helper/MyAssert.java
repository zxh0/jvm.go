package helper;

public class MyAssert {

    public static void assertEquals(Object expected, Object actual) {
        if (!expected.equals(actual)) {
            throw new AssertionError(actual.toString() + " != " + expected.toString());
        }
    }

}
