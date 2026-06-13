import java.util.*;

class RandomizedSet {

    private final List<Integer> values;
    private final Map<Integer, Integer> valueToIndex;
    private final Random random;

    public RandomizedSet() {
        values = new ArrayList<>();
        valueToIndex = new HashMap<>();
        random = new Random();
    }

    public boolean insert(int val) {
        if (valueToIndex.containsKey(val)) {
            return false;
        }

        values.add(val);
        valueToIndex.put(val, values.size() - 1);

        return true;
    }

    public boolean remove(int val) {
        Integer index = valueToIndex.get(val);

        if (index == null) {
            return false;
        }

        int lastIndex = values.size() - 1;
        int lastValue = values.get(lastIndex);

        values.set(index, lastValue);
        valueToIndex.put(lastValue, index);

        values.remove(lastIndex);
        valueToIndex.remove(val);

        return true;
    }

    public int getRandom() {
        int randomIndex = random.nextInt(values.size());
        return values.get(randomIndex);
    }
}