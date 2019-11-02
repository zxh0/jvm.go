package helper;

public class MyAssert {

    public static void assertTrue(String msg, boolean x) {
        if (!x) {
            throw new AssertionError(":(");
        }
    }

    public static void assertTrue(boolean x) {
        if (!x) {
            throw new AssertionError(":(");
        }
    }

    public static void assertFalse(String msg, boolean x) {
        if (x) {
            throw new AssertionError(":(");
        }
    }

    public static void assertFalse(boolean x) {
        if (x) {
            throw new AssertionError(":(");
        }
    }

    public static void assertNull(Object o) {
        if (o != null) {
            throw new AssertionError(":(");
        }
    }

    public static void assertNotNull(Object o) {
        if (o == null) {
            throw new AssertionError(":(");
        }
    }

    public static void assertSame(Object expected, Object actual) {
        if (expected != actual) {
            throw new AssertionError(actual + " != " + expected);
        }
    }

    public static void assertNotSame(Object expected, Object actual) {
        if (expected == actual) {
            throw new AssertionError(actual + " != " + expected);
        }
    }

    public static void assertEquals(int expected, int actual) {
        if (actual != expected) {
            throw new AssertionError(actual + " != " + expected);
        }
    }
    public static void assertEquals(long expected, long actual) {
        if (actual != expected) {
            throw new AssertionError(actual + " != " + expected);
        }
    }
    public static void assertEquals(float expected, float actual) {
        if (actual != expected) {
            throw new AssertionError(actual + " != " + expected);
        }
    }
    public static void assertEquals(double expected, double actual) {
        if (actual != expected) {
            throw new AssertionError(actual + " != " + expected);
        }
    }
    public static void assertEquals(Object expected, Object actual) {
        if (!expected.equals(actual)) {
            throw new AssertionError(actual + " != " + expected);
        }
    }

    public static void assertArrayEquals(Object[] expected, Object[] actual) {
        assertEquals(expected.length, actual.length);
        for (int i = 0; i < expected.length; i++) {
            assertEquals(expected[i], actual[i]);
        }
    }

    public static void fail() {
        throw new RuntimeException(":(");
    }

}
