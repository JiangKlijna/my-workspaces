
import java.lang.reflect.Field;
import java.util.*;

/**
 * 可以把Object当作Map去处理
 * 不支持remove clear等操作
 *
 * @author caojiang
 */
public class ObjectMap extends AbstractMap<String, Object> {

    private final Object object;
    private Set<Entry<String, Object>> set;


    public static ObjectMap convert(Object object) {
        return new ObjectMap(object);
    }

    public ObjectMap(Object object) {
        this.object = object;
    }

    @Override
    public Object put(String key, Object value) {
        try {
            Field field = object.getClass().getDeclaredField(key);
            field.setAccessible(true);
            Object pre = field.get(object);
            field.set(object, value);
            return pre;
        } catch (Exception e) {
            throw new RuntimeException(e);
        }
    }

    @Override
    public void clear() {
        throw new UnsupportedOperationException();
    }

    @Override
    public Object remove(Object key) {
        throw new UnsupportedOperationException();
    }

    @Override
    public Set<Entry<String, Object>> entrySet() {
        if (set == null) synchronized (this) {
            if (set == null) set = newEntrySet();
        }
        return set;
    }

    public Entry<String, Object>[] newEntryArray() {
        Field[] fs = object.getClass().getDeclaredFields();
        ObjectEntry[] oes = new ObjectEntry[fs.length];
        for (int i = 0; i < fs.length; i++) {
            oes[i] = new ObjectEntry(fs[i]);
        }
        return oes;
    }

    public List<Entry<String, Object>> newEntryList() {
        return Arrays.asList(newEntryArray());
    }

    public Set<Entry<String, Object>> newEntrySet() {
        return new ArraySet<>(newEntryArray());
    }

    // 专属Entry
    private class ObjectEntry implements Map.Entry<String, Object> {
        private final Field field;

        public ObjectEntry(Field field) {
            field.setAccessible(true);
            this.field = field;
        }

        @Override
        public String getKey() {
            return field.getName();
        }

        @Override
        public Object getValue() {
            try {
                return field.get(object);
            } catch (IllegalAccessException e) {
                throw new RuntimeException(e);
            }
        }

        @Override
        public Object setValue(Object value) {
            try {
                Object pre = getValue();
                field.set(object, value);
                return pre;
            } catch (IllegalAccessException e) {
                throw new RuntimeException(e);
            }
        }

        @Override
        public boolean equals(Object obj) {
            if (obj == null) return false;
            if (!(obj instanceof ObjectEntry)) return false;
            return field.equals(((ObjectEntry) obj).field);
        }

        @Override
        public int hashCode() {
            return field.hashCode();
        }
    }

    public static class ArraySet<T> extends AbstractSet<T> {
        public final List<T> set;

        public ArraySet(T[] set) {
            this.set = Arrays.asList(set);
        }

        @Override
        public boolean remove(Object o) {
            throw new UnsupportedOperationException();
        }

        @Override
        public Iterator<T> iterator() {
            return set.iterator();
        }

        @Override
        public int size() {
            return set.size();
        }
    }
}
