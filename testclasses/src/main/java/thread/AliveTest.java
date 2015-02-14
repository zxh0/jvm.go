package thread;

public class AliveTest {
    
    public static void main(String[] args) throws InterruptedException {
        Thread t = new Thread() {

            @Override
            public void run() {
                try {
                    Thread.sleep(2000);
                } catch (InterruptedException e) {
                    e.printStackTrace(System.err);
                }
            }
            
        };
        
        if (t.isAlive() != false) {
            System.out.println("t is alive!");
            return;
        }
        
        t.start();
        Thread.sleep(1000);
        if (t.isAlive() != true) {
            System.out.println("t is not alive!");
            return;
        }
        
        t.join();
        if (t.isAlive() != false) {
            System.out.println("t is not dead!");
            return;
        }
        
        System.out.println("OK!");
    }
    
}
