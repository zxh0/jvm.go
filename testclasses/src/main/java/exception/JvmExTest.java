package exception;

public class JvmExTest {
    
    public static void main(String[] args) {
        try {
            // NPE
            //arraylength(null);
            //athrow(null);
            
            // ClassCastException
            checkcast();
        } catch (Exception e) {
            e.printStackTrace(System.err);
        }
    }
    
    static void arraylength(int[] x) {
        int y = x.length;
    }
    
    static void athrow(Exception ex) throws Exception {
        throw ex;
    }
    
    static void checkcast() {
        Object x = "String";
        Integer y = (Integer) x;
    }
    
//    private static void test() {
//        new NPETest().foo();
//    }
//    
//    private void foo() {
//        bar();
//    }
//    
//    private void bar() {
//        Object x = null;
//        synchronized(x) {
//            System.out.println("BAD!");
//        }
//    }
    
}
