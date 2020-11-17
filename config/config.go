package config

import (
	"time"

	"github.com/spf13/pflag"

	"github.com/wflysnow/my-middleware-operator/pkg/apis/redis/v1alpha1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	rl "k8s.io/client-go/tools/leaderelection/resourcelock"
)

const (
	DefaultLeaseDuration = 15 * time.Second
	DefalutRenewDeadlind = 10 * time.Second
	DefaultRetryPeriod   = 2 * time.Second
)

func DefaultLeaderElectionConfiguration() v1alpha1.LeaderElectionConfiguration {
	return v1alpha1.LeaderElectionConfiguration{
		LeaderElect:   false,
		LeaseDuration: metav1.Duration{Duration: DefaultLeaseDuration},
		RenewDeadline: metav1.Duration{Duration: DefalutRenewDeadlind},
		RetryPeriod:   metav1.Duration{Duration: DefaultRetryPeriod},
		ResourceLock:  rl.EndpointsResourceLock,
	}
}

func BindFlags(l *v1alpha1.LeaderElectionConfiguration, fs *pflag.FlagSet) {
	fs.BoolVar(&l.LeaderElect, "leader-elect", l.LeaderElect, ""+
		"start a leader election client and gain leadership before "+
		"executing the main loop. Enable this when running replicated "+
		"components for high availability.")
	fs.DurationVar(&l.LeaseDuration.Duration, "leader-elect-lease-duration", l.LeaseDuration.Duration, ""+
		"The duration that non-leader candidates will wait after observing a leadership "+
		"renewal until attempting to acquire leadership of a led but unrenewed leader "+
		"slot. This is effectively the maximum duration that a leader can be stopped "+
		"before it is replaced by another candidate. This is only applicable if leader "+
		"election is enabled.")
	fs.DurationVar(&l.RenewDeadline.Duration, "leader-elect-renew-deadline", l.RenewDeadline.Duration, ""+
		"The interval between attempts by the acting master to renew a leadership slot "+
		"before it stops leading. This must be less than or equal to the lease duration. "+
		"This is only applicable if leader election is enabled.")
}
