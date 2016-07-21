def repo = "permissions"
def dockerUser = "discoenv"

node {
    stage "Build"
    checkout scm

    sh 'git rev-parse HEAD > GIT_COMMIT'
    git_commit = readFile('GIT_COMMIT').trim()
    echo git_commit

    dockerRepo = "${dockerUser}/${repo}:${env.BRANCH_NAME}"

    sh "docker build --rm --build-arg git_commit=${git_commit} -t ${dockerRepo} ."


    stage "Test"
    sh "docker run --rm -w /go/src/github.com/cyverse-de/${repo} --entrypoint 'gb' ${dockerRepo} test"

    stage "Docker Push"
    sh "docker push ${dockerRepo}"
}
