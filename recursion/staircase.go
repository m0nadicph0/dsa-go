package recursion

/*
You are given a staircase with 'n' steps, and you can climb either 1 step, 2 steps, or 3 steps at a time.
How many different ways can you reach the top of the staircase?

For example, if the staircase has 4 steps, there are 7 different ways to reach the top:

1 step + 1 step + 1 step + 1 step
1 step + 1 step + 2 steps
1 step + 2 steps + 1 step
2 steps + 1 step + 1 step
2 steps + 2 steps
1 step + 3 steps
3 steps + 1 step
The problem can be solved using recursion. The base cases would be:

If there are 0 steps left, there's only one way to "climb" it, which is by doing nothing.
If there's 1 step left, there's only one way to climb it, which is by taking one step.
The recursive relation can be defined as follows:

The number of ways to climb 'n' steps is equal to the sum of the number of ways to climb 'n-1' steps, 'n-2' steps, and 'n-3' steps.
*/

func CountWaysToClimbStairs(n int) int {
	switch n {
	case 0, 1:
		return 1
	case 2:
		return 2
	default:
		return CountWaysToClimbStairs(n-1) + CountWaysToClimbStairs(n-2) + CountWaysToClimbStairs(n-3)
	}
}
