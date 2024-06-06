package problem

type IProblem interface {
	ProblemId() string
	ContestName() string
	Rank() string

	ProblemDirPath() string

	ProblemUrl() string
	SubmissionUrl() string
}
