package thread;

public class RunnableTest implements Runnable {
    
    public static void main(String[] args) {
        Thread t = new Thread(new RunnableTest());
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
