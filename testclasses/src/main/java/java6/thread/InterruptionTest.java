package java6.thread;

public class InterruptionTest {
    
    public static void main(String[] args) throws InterruptedException {
        Thread t = new Thread(new Runnable() {

            @Override
            public void run() {
                try {
                    Thread.sleep(1000);
                } catch (InterruptedException ex) {
                    System.out.println(ex.getMessage());
                }
            }
            
        });
        t.start();
        
        Thread.sleep(500);
        t.interrupt();
    }
    
}
