class RandomizedSet {
    constructor() {
        this.values = [];
        this.valueToIndex = new Map();
    }

    insert(val) {
        if (this.valueToIndex.has(val)) {
            return false;
        }

        this.values.push(val);
        this.valueToIndex.set(val, this.values.length - 1);

        return true;
    }

    remove(val) {
        if (!this.valueToIndex.has(val)) {
            return false;
        }

        const index = this.valueToIndex.get(val);
        const lastIndex = this.values.length - 1;
        const lastValue = this.values[lastIndex];

        this.values[index] = lastValue;
        this.valueToIndex.set(lastValue, index);

        this.values.pop();
        this.valueToIndex.delete(val);

        return true;
    }

    getRandom() {
        const randomIndex = Math.floor(
            Math.random() * this.values.length
        );

        return this.values[randomIndex];
    }
}