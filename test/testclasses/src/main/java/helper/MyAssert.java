package helper;

public class MyAssert {

    public static void assertTrue(String msg, boolean x) {
        if (!x) {
            throw new AssertionError("TODO");
        }
    }

    public static void assertTrue(boolean x) {
        if (!x) {
            throw new AssertionError("TODO");
        }
    }

    public static void assertFalse(String msg, boolean x) {
        if (x) {
            throw new AssertionError("TODO");
        }
    }

    public static void assertFalse(boolean x) {
        if (x) {
            throw new AssertionError("TODO");
        }
    }

    public static void assertNull(Object o) {
        if (o != null) {
            throw new AssertionError("TODO");
        }
    }

    public static void assertNotNull(Object o) {
        if (o == null) {
            throw new AssertionError("TODO");
        }
    }

    public static void assertSame(Object expected, Object actual) {
        if (expected != actual) {
            throw new AssertionError(actual.toString() + " != " + expected.toString());
        }
    }

    public static void assertNotSame(Object expected, Object actual) {
        if (expected == actual) {
            throw new AssertionError(actual.toString() + " != " + expected.toString());
        }
    }

    public static void assertEquals(Object expected, Object actual) {
        if (!expected.equals(actual)) {
            throw new AssertionError(actual.toString() + " != " + expected.toString());
        }
    }

    public static void assertEquals(String msg, Object expected, Object actual) {
        if (!expected.equals(actual)) {
            throw new AssertionError(actual.toString() + " != " + expected.toString());
        }
    }

    public static void fail() {
        throw new RuntimeException("TODO");
    }

}
