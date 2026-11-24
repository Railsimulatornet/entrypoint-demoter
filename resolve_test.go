package entrypoint_demoter

import "testing"

func TestResolveIdsRejectsUnusableIds(t *testing.T) {
	for _, field := range []string{"UID", "GID"} {
		for _, value := range []string{"-1", "-4294967296", "4294967295", "4294967296", "8589934592", "18446744073709551616", "not-a-number"} {
			t.Run(field+"/"+value, func(t *testing.T) {
				t.Setenv("UID", "0")
				t.Setenv("GID", "0")
				t.Setenv(field, value)
				uid, gid, err := ResolveIds("")
				if err == nil {
					t.Fatalf("unusable %s was accepted as uid=%d gid=%d", field, uid, gid)
				}
			})
		}
	}
}

func TestResolveIdsPreservesValidIds(t *testing.T) {
	for _, test := range []struct {
		value string
		want  uint32
	}{
		{"", 0}, {"0", 0}, {"1000", 1000}, {"+1000", 1000}, {"001000", 1000}, {"65534", 65534},
		{"2147483648", 2147483648}, {"4294967294", 4294967294},
	} {
		t.Run(test.value, func(t *testing.T) {
			t.Setenv("UID", test.value)
			t.Setenv("GID", test.value)
			uid, gid, err := ResolveIds("")
			if err != nil || uid != test.want || gid != test.want {
				t.Fatalf("ResolveIds() = %d, %d, %v; want %d", uid, gid, err, test.want)
			}
		})
	}
}
