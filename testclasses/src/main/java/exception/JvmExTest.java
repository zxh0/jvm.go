package exception;

public class JvmExTest {
    
    public static void main(String[] args) {
        try {
            //arraylength(null);
            athrow(null);
        } catch (Exception e) {
            e.printStackTrace(System.err);
        }
    }
    
    static void arraylength(int[] x) {
        if (x.length > 0) {}
    }
    
    static void athrow(Exception ex) throws Exception {
        throw ex;
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
