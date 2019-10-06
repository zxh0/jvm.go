package java8;

import junit.framework.Assert;
import libs.junit.UnitTestRunner;
import org.junit.Test;

/**
 * Created with IntelliJ IDEA.
 * User: beyond
 * Email: beyondblog@outlook.com
 * Date: 15/4/6
 * Time: 下午5:28
 */
public class InterfaceDefaultMethodTest {

    class TestInterfaceDefaultTest extends FirstTestClass implements FirstTest, SecondTest {
    
    }

    class TestInterfaceFirstTestClass implements FirstTest, SecondTest, ThirdTest, DefaultTest {
    
    }

    class FirstTestClass implements ThirdTest, FirstTest {
    
    }

    interface FirstTest extends DefaultTest {
    
    }

    interface DefaultTest {
        default String test() {
            return "DefaultTest";
        }
    }

    interface SecondTest extends FirstTest {
        default String test() {
            return "SecondTest";
        }
    }

    interface ThirdTest extends DefaultTest, SecondTest {
        default String test() {
            return "ThirdTest";
        }
    }

    public static void main(String[] args) {
        UnitTestRunner.run(InterfaceDefaultMethodTest.class);
    }

    //@Test
    public void Test1() {
        TestInterfaceDefaultTest testInterfaceDefaultTest = new TestInterfaceDefaultTest();
        DefaultTest defaultTest = testInterfaceDefaultTest;
        FirstTestClass firstTestClass = testInterfaceDefaultTest;
        SecondTest secondTest = testInterfaceDefaultTest;
        ThirdTest thirdTest = testInterfaceDefaultTest;
        Assert.assertEquals(defaultTest.test(), "ThirdTest");
        Assert.assertEquals(firstTestClass.test(), "ThirdTest");
        Assert.assertEquals(secondTest.test(), "ThirdTest");
        Assert.assertEquals(thirdTest.test(), "ThirdTest");
    }

    //@Test
    public void Test2() {
        TestInterfaceFirstTestClass testInterfaceFirstTestClass = new TestInterfaceFirstTestClass();
        DefaultTest defaultTest = testInterfaceFirstTestClass;
        FirstTest firstTest = testInterfaceFirstTestClass;
        SecondTest secondTest = testInterfaceFirstTestClass;
        ThirdTest thirdTest = testInterfaceFirstTestClass;

//        System.out.println(defaultTest.test());
//        System.out.println(firstTest.test());
//        System.out.println(secondTest.test());
//        System.out.println(thirdTest.test());
        Assert.assertEquals(defaultTest.test(), "ThirdTest");
        Assert.assertEquals(firstTest.test(), "ThirdTest");
        Assert.assertEquals(secondTest.test(), "ThirdTest");
        Assert.assertEquals(thirdTest.test(), "ThirdTest");
    }
}
