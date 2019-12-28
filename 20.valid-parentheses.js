/*
 * @lc app=leetcode id=20 lang=javascript
 *
 * [20] Valid Parentheses
 *
 * https://leetcode.com/problems/valid-parentheses/description/
 *
 * algorithms
 * Easy (37.72%)
 * Likes:    3803
 * Dislikes: 190
 * Total Accepted:    792.8K
 * Total Submissions: 2.1M
 * Testcase Example:  '"()"'
 *
 * Given a string containing just the characters '(', ')', '{', '}', '[' and
 * ']', determine if the input string is valid.
 * 
 * An input string is valid if:
 * 
 * 
 * Open brackets must be closed by the same type of brackets.
 * Open brackets must be closed in the correct order.
 * 
 * 
 * Note that an empty string is also considered valid.
 * 
 * Example 1:
 * 
 * 
 * Input: "()"
 * Output: true
 * 
 * 
 * Example 2:
 * 
 * 
 * Input: "()[]{}"
 * Output: true
 * 
 * 
 * Example 3:
 * 
 * 
 * Input: "(]"
 * Output: false
 * 
 * 
 * Example 4:
 * 
 * 
 * Input: "([)]"
 * Output: false
 * 
 * 
 * Example 5:
 * 
 * 
 * Input: "{[]}"
 * Output: true
 * 
 * 
 */

// @lc code=start
/**
 * @param {string} s
 * @return {boolean}
 */
var isValid = function (s) {


    // This Stack is written using the pseudoclassical pattern

    // Creates a stack
    var Stack = function () {
        this.count = 0;
        this.storage = {};
    }

    // Adds a value onto the end of the stack
    Stack.prototype.push = function (value) {
        this.storage[this.count] = value;
        this.count++;
    }

    // Removes and returns the value at the end of the stack
    Stack.prototype.pop = function () {
        // Check to see if the stack is empty
        if (this.count === 0) {
            return undefined;
        }

        this.count--;
        var result = this.storage[this.count];
        delete this.storage[this.count];
        return result;
    }

    // Returns the length of the stack
    Stack.prototype.size = function () {
        return this.count;
    }

    for (var element of s) {
        switch (element) {
            case '(':
                Stack.push(')');
                break;

            case '{':
                Stack.push('}');
                break;

            case '[':
                Stack.push(']');
                break;

            case ')':
            case '}':
            case ']':
                var cur = Stack.pop();
                if (cur != element) {
                    return false;
                }
                break;

            default:
                break;
        }
    }

    return true;
};
// @lc code=end