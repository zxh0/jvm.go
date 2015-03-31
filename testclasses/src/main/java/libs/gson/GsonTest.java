package libs.gson;

import com.google.gson.Gson;
import junit.framework.Assert;
import org.junit.Test;
import sun.reflect.generics.factory.CoreReflectionFactory;
import sun.reflect.generics.factory.GenericsFactory;
import sun.reflect.generics.repository.FieldRepository;
import sun.reflect.generics.scope.ClassScope;
import sun.reflect.generics.tree.TypeSignature;
import sun.reflect.generics.tree.TypeVariableSignature;
import sun.reflect.generics.visitor.Reifier;

import java.lang.reflect.Field;

class MyFieldRepository extends FieldRepository {
    protected MyFieldRepository(String var1, GenericsFactory var2) {

        super(var1, var2);
    }


    public static MyFieldRepository make(String var0, GenericsFactory var1) {
        return new MyFieldRepository(var0, var1);
    }

    public TypeSignature getTree2() {
        return ((TypeSignature) this.getTree());
    }

    public Reifier getReifier2() {
        return this.getReifier();

    }
}


public class GsonTest {

    public int x;
    public float y;
    public String z;

    // Accessor for factory
    private static GenericsFactory getFactory() throws Exception {
        JsonTest<String> jsonTest = new JsonTest<String>();
        Field field = jsonTest.getClass().getDeclaredField("x");
        Class<?> c = field.getDeclaringClass();
        // create scope and factory
        return CoreReflectionFactory.make(c, ClassScope.make(c));
    }

    // Accessor for generic info repository
    private static MyFieldRepository getGenericInfo() throws Exception {

        return MyFieldRepository.make("TT;", getFactory()); //return cached repository
    }

    public static void main(String[] args) throws Exception {
        // UnitTestRunner.run(GsonTest.class);
        MyFieldRepository fieldRepository = getGenericInfo();

        Reifier var1 = fieldRepository.getReifier2();
        TypeSignature tree = ((TypeSignature) fieldRepository.getTree2());

        TypeVariableSignature typeVariableSignature = (TypeVariableSignature) tree;

        System.out.println(typeVariableSignature.getIdentifier()); //T

        typeVariableSignature.accept(var1); //here is a bug
//        tree.accept(var1);
        System.out.println(var1.getResult()); //null
        System.out.println(fieldRepository.getGenericType()); //null


    }

    @Test
    public void gson() {
        GsonTest obj = new GsonTest();
        obj.x = 100;
        obj.y = 3.14f;
        obj.z = "hello";
        Assert.assertEquals("{\"x\":100,\"y\":3.14,\"z\":\"hello\"}", new Gson().toJson(obj));
    }

    @Test
    public void testGenericsType() {
        try {
            JsonTest<String> jsonTest = new JsonTest<String>();
            jsonTest.x = "beyond";
            System.out.println(new Gson().toJson(jsonTest));
            Assert.assertEquals("{\"x\":\"beyond\"}", new Gson().toJson(jsonTest));
            System.out.println("end");
        } catch (Exception e) {
            e.printStackTrace();
        }
    }

    @Test
    public void TestGenerics() {
        try {
        } catch (Exception e) {
            e.printStackTrace();
        }
    }
}

class JsonTest<T> {
    public T x;
}
