package reflection;

public class DeclaringClassTest {
    
    static class A {
        static class B {
            class C {
                
            }
        }
    }
    
    public static void main(String[] args) {
        System.out.println(DeclaringClassTest.class.getDeclaringClass());
        System.out.println(A.class.getDeclaringClass());
        System.out.println(A.B.class.getDeclaringClass());
        System.out.println(A.B.C.class.getDeclaringClass());
    }
    
}
