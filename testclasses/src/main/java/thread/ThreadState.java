package thread;

public class ThreadState {
    
    public static void main(String[] args) {
        Thread t = new Thread();
        if (t.isAlive() != false) {
            System.out.println("t is alive!");
        }
        
        System.out.println("OK!");
    }
    
}
