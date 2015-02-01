public class ThreadTest {
    
    public static void main(String[] args) {
        Thread t = new Thread();//.start();
        System.out.println(t.getThreadGroup().getName());
    }
    
}
