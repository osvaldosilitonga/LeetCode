import solution.two_sum.TwoSum;

import java.util.Arrays;

public class Main {
    public static void main(String[] args) {

        System.out.println("Hello world!");

        TwoSum twoSumObj = new TwoSum();

        int[] nums = {2,7,11,15};
        int target = 9;

        int[] result = twoSumObj.twoSum(nums, target); // result must be = [0, 1]

        System.out.println("Result : " + Arrays.toString(result));
    }
}