mockgen --source=../internal/repository/repository.go --destination=../internal/repository/mockgen_generated/mock_repository/mock_repository.go


#reference: https://budimanokky93.medium.com/golang-unit-test-mocking-technique-4d9225d1dc76
#install this first:
#https://github.com/golang/mock
#release mockgen tools and mock go module
#then run this .sh
#run go get all && go mod tidy && go mod vendor