package thread;

public class DaemonTest {
    
    public static void main(String[] args) {
        Thread mainThread = Thread.currentThread();
        System.out.println("main thread is daemon: " + mainThread.isDaemon());
        
        Thread newThread = new Thread();
        System.out.println("new thread is daemon: " + newThread.isDaemon());
        
        newThread.setDaemon(true);
        System.out.println("new thread is daemon: " + newThread.isDaemon());
    }
    
}
