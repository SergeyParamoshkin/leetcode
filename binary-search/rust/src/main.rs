pub fn search(nums: Vec<i32>, target: i32) -> i32 {
    let mut pivot = 0 as usize;
    let mut left = 0 as usize;
    let mut right = nums.len() as usize;
    while left < right {
        pivot = left + (right - left) / 2;
        if nums[pivot] == target {
            return pivot as i32;
        }
        if target < nums[pivot] {
            right = pivot - 1;
        } else {
            left = pivot + 1;
        }
    }

    return -1;
}

fn main() {
    let nums = vec![-1, 0, 3, 5, 9, 12];
    let target = 9;
    search(nums, target);
}
