package solution.two_sum;
public class TwoSum {
    public int[] twoSum(int[] nums, int target) {
        for (int i = 1; i < nums.length; i++) {
            for (int x = i; x < nums.length; x++) {
                if (nums[x]+nums[x-i] == target) {
                    return new int[] {x-i, x};
                }
            }
        }
        return null;
    }
}
