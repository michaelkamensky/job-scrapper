/*
==========
Cariddi
==========

This program is free software: you can redistribute it and/or modify
it under the terms of the GNU General Public License as published by
the Free Software Foundation, either version 3 of the License, or
(at your option) any later version.

This program is distributed in the hope that it will be useful,
but WITHOUT ANY WARRANTY; without even the implied warranty of
MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
GNU General Public License for more details.

You should have received a copy of the GNU General Public License
along with this program.  If not, see http://www.gnu.org/licenses/.

	@Repository:  https://github.com/edoardottt/cariddi

	@Author:      edoardottt, https://edoardottt.com

	@License: https://github.com/edoardottt/cariddi/blob/main/LICENSE

*/

package crawler

import (
	"fmt"
	"strings"

	urlUtils "github.com/edoardottt/cariddi/internal/url"
)

// RetrieveBody retrieves the body (in the response) of a url.
func RetrieveBody(target *string) string {
	sb, err := GetRequest(*target)
	if err == nil && sb != "" {
		return sb
	}

	return ""
}

// IgnoreMatch checks if the URL should be ignored or not.
func IgnoreMatch(url string, ignoreSlice *[]string) bool {
	for _, ignore := range *ignoreSlice {
		if strings.Contains(url, ignore) {
			return true
		}
	}

	return false
}

// intensiveOk checks if a given url can be crawled
// in intensive mode (if the 2nd level domain matches with
// the inputted target). If maxSubdomainDepth > 0, it limits
// how many subdomain labels are allowed.
func intensiveOk(target string, urlInput string, debug bool, maxSubdomainDepth int) bool {
	root, err := urlUtils.GetRootHost(urlInput)
	if err != nil {
		if debug {
			fmt.Println(err.Error() + ": " + urlInput)
		}

		return false
	}

	if root != target {
		return false
	}

	if maxSubdomainDepth <= 0 {
		return true
	}

	host := urlUtils.GetHost(urlInput)
	if host == "" {
		if debug {
			fmt.Println("unable to parse host: " + urlInput)
		}
		return false
	}

	host = urlUtils.RemovePort(host)
	depth, ok := subdomainDepth(host, target)
	if !ok {
		if debug {
			fmt.Println("unable to determine subdomain depth: " + urlInput)
		}
		return false
	}

	return depth <= maxSubdomainDepth
}

// subdomainDepth returns the number of labels before the root domain.
// Example: host "a.b.example.com" with root "example.com" => depth 2.
func subdomainDepth(host string, root string) (int, bool) {
	if host == root {
		return 0, true
	}

	if !strings.HasSuffix(host, "."+root) {
		return 0, false
	}

	hostParts := strings.Split(host, ".")
	rootParts := strings.Split(root, ".")
	if len(hostParts) < len(rootParts) {
		return 0, false
	}

	return len(hostParts) - len(rootParts), true
}
