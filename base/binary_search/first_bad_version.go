/**
 * Forward declaration of isBadVersion API.
 * @param   version   your guess about first bad version
 * @return 	 	      true if current version is bad
 *			          false if current version is good
 * func isBadVersion(version int) bool;
 */

func firstBadVersion(n int) int {
	return search(0, n)
}

func search(start, end int) int {
	endBad := isBadVersion(end)
	d := end - start
	if d <= 1 {
		return end
	}

	pre := (end-start)/2 + start
	preBad := isBadVersion(pre)
	if endBad && preBad {
		return search(start, pre)
	}
	if endBad && !preBad {
		return search(pre, end)
	}

	return end
}
