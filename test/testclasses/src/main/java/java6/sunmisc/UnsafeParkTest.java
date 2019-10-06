package java6.sunmisc;

import libs.junit.UnitTestRunner;
import org.junit.Assert;
import org.junit.Test;

import java.util.Date;
import java.util.concurrent.locks.LockSupport;
import java.util.concurrent.locks.ReentrantLock;


/**
 * Created by beyond on 15-3-27.
 */
public class UnsafeParkTest {
    public static void main(String[] args) throws Exception {
        UnitTestRunner.run(UnsafeParkTest.class);
    }

    /**
     * park is released by unpark occurring after park
     */
    @Test
    public void testPark() throws Exception {
        Thread t = new Thread(new Runnable() {
            public void run() {
                try {
                    LockSupport.park();
                } catch (Exception e) {
                    throw e;
                }
            }
        });
        try {
            t.start();
            Thread.sleep(1000);
            LockSupport.unpark(t);
            t.join();
        } catch (Exception e) {
            throw e;
        }
    }

    /**
     * park is released by unpark occurring before park
     */
    @Test
    public void testPark2() throws Exception {
        Thread t = new Thread(new Runnable() {
            public void run() {
                try {
                    Thread.sleep(1000);
                    LockSupport.park();
                } catch (Exception e) {
                    System.out.println(e.getMessage());
                }
            }
        });
        try {
            t.start();
            LockSupport.unpark(t);
            t.join();
        } catch (Exception e) {
            throw e;
        }
    }

    /**
     * park is released by interrupt
     */
    @Test
    public void testPark3() {
        Thread t = new Thread(new Runnable() {
            public void run() {
                try {
                    //Thread.sleep(2000);
                    LockSupport.park();
                    //TODO
                    Assert.assertTrue(Thread.interrupted());
                } catch (Exception e) {
                    Assert.assertTrue(Thread.interrupted());
                    System.out.println(e.getMessage());
                }
            }
        });
        try {
            t.start();
            Thread.sleep(1000);
            t.interrupt();
            t.join();
        } catch (Exception e) {
            System.out.println(e.getMessage());
        }
    }

    /**
     * park returns if interrupted before park
     */
    @Test
    public void testPark4() {
        final ReentrantLock lock = new ReentrantLock();
        lock.lock();
        Thread t = new Thread(new Runnable() {
            public void run() {
                try {
                    lock.lock();
                    LockSupport.park();
                } catch (Exception e) {
                    System.out.println(e.getMessage());

                }
            }
        });
        try {
            t.start();
            t.interrupt();
            lock.unlock();
            t.join();
        } catch (Exception e) {
            System.out.println(e.getMessage());
        }
    }

    /**
     * parkNanos times out if not unparked
     */
    @Test
    public void testParkNanos() {
        Thread t = new Thread(new Runnable() {
            public void run() {
                try {
                    LockSupport.parkNanos(1000);
                } catch (Exception e) {
                    System.out.println(e.getMessage());
                }
            }
        });
        try {
            t.start();
            t.join();
        } catch (Exception e) {
            System.out.println(e.getMessage());
        }
    }


    /**
     * parkUntil times out if not unparked
     */
    @Test
    public void testParkUntil() {
        Thread t = new Thread(new Runnable() {
            public void run() {
                try {
                    long d = new Date().getTime() + 100;
                    LockSupport.parkUntil(d);
                } catch (Exception e) {
                    System.out.println(e.getMessage());
                }
            }
        });
        try {
            t.start();
            t.join();
        } catch (Exception e) {
            System.out.println(e.getMessage());
        }
    }
}
