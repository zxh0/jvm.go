package thread;

public class ThreadTest1 implements Runnable {
    
    public static void main(String[] args) {
        Thread t = new Thread(new ThreadTest1());
        t.start();
        
        for (int i = 0; i < 100; i++) {
            System.out.println("main:" + i);
        }
    }

    @Override
    public void run() {
        for (int i = 0; i < 100; i++) {
            System.out.println("run:" + i);
        }
    }
    
}
