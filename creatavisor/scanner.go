package creatavisor

import (
	"bufio"
	"regexp"
)

// Trim off whitespace around the info - match least greedy, grab as much space on both sides

//  fmt.Sprintf("UPGRADE \"%s\" NEEDED at %s: %s", plan.Name, plan.DueAt(), plan.Info)
//    if !p.Time.IsZero() {
//      return fmt.Sprintf("time: %s", p.Time.UTC().Format(time.RFC3339))
//    }
//    return fmt.Sprintf("height: %d", p.Height)
var upgradeRegex = regexp.MustCompile(`UPGRADE "(.*)" NEEDED at ((height): (\d+)|(time): (\S+)):\s+(\S*)`)

// UpgradeInfo is the details from the regexp
type UpgradeInfo struct {
	Name string
	Info string
}

// WaitForUpdate will listen to the scanner until a line matches upgradeRegexp.
// It returns (info, nil) on a matching line
// It returns (nil, err) if the input stream errored
// It returns (nil, nil) if the input closed without ever matching the regexp
func WaitForUpdate(scanner *bufio.Scanner) (*UpgradeInfo, error) {
	for scanner.Scan() {
		line := scanner.Text()
		if upgradeRegex.MatchString(line) {
			subs := upgradeRegex.FindStringSubmatch(line)
			info := UpgradeInfo{
				Name: subs[1],
				Info: subs[7],
			}
			return &info, nil
		}
	}
	return nil, scanner.Err()
}
