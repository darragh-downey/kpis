package perf

import "github.com/robatussum/kpis/model"

// AFR
// Annualised Failure Rate the estimated probability of a device or component failing during a full year
// of operation.
func AFR() {}

// MDT
// Mean Downtime is the average time an asset/system is unavailable
// Returns Mean Downtime as seconds
func MDT(c1 model.Component) int64 {
	var mdt int64 = 0

	for _, failure := range c1.Failures {
		mdt += failure.EndTime.Unix() - failure.StartTime.Unix()
	}

	if mdt == 0 || len(c1.Failures) == 0 {
		return int64(0)
	}

	return mdt / int64(len(c1.Failures))
}

// MDTS
// Mean Downtime Serial is the average time for a system with serial components is unavailable
//
// Equivalent to:
// mdt(c1;c2)
// = mtbf(c1) * mdt(c2) + mtbf(c2) * mdt(c1) / mtbf(c1) + mtbf(c2)
//
// Where:
// c1;c2 is the network in which the components are arranged in series
// Returns Mean Downtime as seconds
func MDTS(c1, c2 model.Component) int64 {
	top := MTBF(c1)*MDT(c2) + MTBF(c2)*MDT(c1)
	div := MTBF(c1) + MTBF(c2)

	if div == 0 {
		return 0
	}

	return top / div
}

// MDTP
// Mean Downtime Parallel is the average time for a system with parallel components is unavailable
//
// Equivalent to:
// mdt(c1||c2)
// = mdt(c1) * mdt(c2) / mdt(c1) + mdt(c2)
//
// Where:
// c1||c2 is the network in which the components are arranged in parallel
// Returns Mean Downtime as seconds
func MDTP(c1, c2 model.Component) int64 {
	top := MDT(c1) * MDT(c2)
	div := MDT(c1) + MDT(c2)

	if div == 0 {
		return 0
	}

	return top / div
}

// MTBF
// Mean Time Between Failures metric denotes time between failures in a system which can be repaired,
// it should not be used for systems where repair is not possible.
// Returns Mean Time Between Failures as seconds
func MTBF(c1 model.Component) int64 {
	var mtbf int64 = 0

	// need to test for the case of 0||1 failures
	for i := 0; i < len(c1.Failures)-1; i++ {
		// end of the current failure (uptime) to start of the next failure (downtime)
		mtbf += c1.Failures[i].EndTime.Unix() - c1.Failures[i+1].StartTime.Unix()
	}

	if mtbf == 0 || len(c1.Failures)-1 == 0 {
		return int64(0)
	}

	return mtbf / int64(len(c1.Failures)-1)
}

// MTBFS
// Mean Time Between Failures Series for systems in a network where the components are running in a series
// Equivalent to:
// mtbf(c1;c2)
// = 1 / (1/mtbf(c1) + 1/mtbf(c2))
// = mtbf(c1) * mtbf(c2) / mtbf(c1) + mtbf(c2)
//
// Where:
// c1;c2 is the network in which the components are arranged in series.
// Returns Mean Time Between Failures as seconds
func MTBFS(c1, c2 model.Component) int64 {
	top := MTBF(c1) * MTBF(c2)
	div := MTBF(c1) + MTBF(c2)

	if div == 0 {
		return 0
	}

	return top / div
}

// MTBFP
// Mean Time Between Failures Parallel for systems running in parallel
// Equivalent to:
// mtbf(c1||c2)
// = 1 / (1 / mtbf(c1) * PF(c2, mdt(c1) + 1/mtbf(c2) * PF(c1, mdt(c2))) )
// = mtbf(c1) * mtbf(c2) / mdt(c1) + mdt(c2)
//
// Where:
// c1 || c2 is the network in which the components are arranged in parallel, and,
// PF(c, t) is the probability of failure of component c during "vulnerability window" t.
// Returns Mean Time Between Failures as seconds
func MTBFP(c1, c2 model.Component) int64 {
	top := MTBF(c1) * MTBF(c2)
	div := MDT(c1) + MDT(c2)

	if div == 0 {
		return 0
	}

	return top / div
}

// MTTA
// Mean Time to Acknowledge describes the responsiveness of maintenance teams as it measures the
// time from when the maintenance teams are alerted to the existence of a fault to the time the
// maintenance team acknowledge the fault
// Returns Mean Time to Acknowledge as seconds
func MTTA() int64 {
	return 0
}

// MTTD
// Mean Time to Detection describes the average time for parties to detect and report a failure
// Returns Mean Time to Detect as seconds
func MTTD() int64 {
	return 0
}

// MTTF
// Mean Time to Failure is the average time between non-repairable failures of a system
// Returns Mean Time to Failure as seconds
func MTTF() int64 {
	return 0
}

// MTTRecover
// Mean Time to Recover is a key metric to assist in determining where the problem lies within the
// Maintenance process for a system and is useful for assessing the speed of the overall recovery process
// It covers the period when the system becomes unavailable to the time it becomes available again.
// Returns Mean Time to Recover as seconds
func MTTRecover() int64 {
	return 0
}

// MTTRepair
// Mean Time to Repair represents the average time it takes to repair a system and includes the repair time
// plus any additional testing time. It is important to note that the timing of this process does not conclude
// until the system is fully functional
// Returns Mean Time to Repair as seconds
func MTTRepair() int64 {
	return 0
}

// MTTResponse
// Mean Time to Respond is a metric which illustrates the average time taken to recover from a failure
// starting from the time the team is first alerted to that failure.
// Returns Mean Time to Response as seconds
func MTTResponse() int64 {
	return 0
}

// MTTResolve
// Mean Time to Resolve is the average time taken to resolve a failure in a system entirely. Mean Time to Resolve
// is a key metric which assists in improving customer satisfaction as it covers the time spent detecting the failure,
// diagnosis, repairing the fault and assurance time
// (assurance time covers the time it takes to ensure the issue won't repeat)
// Returns Mean Time to Resolve as seconds
func MTTResolve() int64 {
	return 0
}

// Uptime is a useful metric for determining the availability of a system
// Returns Uptime as seconds
func Uptime(c1 model.Component) int64 {
	var up int64 = 0

	for i := 0; i < len(c1.Availability)-1; i++ {
		up += c1.Availability[i].StartTime.Unix() - c1.Availability[i+1].EndTime.Unix()
	}

	return up
}
