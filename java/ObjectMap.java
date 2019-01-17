
import java.lang.reflect.Field;
import java.util.AbstractMap;
import java.util.LinkedHashSet;
import java.util.Map;
import java.util.Set;


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
    public synchronized Set<Entry<String, Object>> entrySet() {
        if (set == null) {
            set = new LinkedHashSet<>();
            initEntrySet();
        }
        return set;
    }

    private void initEntrySet() {
        for (Field field : object.getClass().getDeclaredFields()) {
            set.add(new ObjectEntry(field));
        }
    }

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
}
