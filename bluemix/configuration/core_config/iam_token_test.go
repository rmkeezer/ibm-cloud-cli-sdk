package core_config

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var TestIAMTokenData = []string{
	//token from password
	"eyJraWQiOiIyMDE3MTAzMC0wMDowMDowMCIsImFsZyI6IlJTMjU2In0.eyJpYW1faWQiOiJJQk1pZC0yNzAwMDZWOEhNIiwiaWQiOiJJQk1pZC0yNzAwMDZWOEhNIiwicmVhbG1pZCI6IklCTWlkIiwiaWRlbnRpZmllciI6IjI3MDAwNlY4SE0iLCJnaXZlbl9uYW1lIjoiT0UgUnVudGltZXMiLCJmYW1pbHlfbmFtZSI6IlN5c3RlbSBVc2VyIiwibmFtZSI6Ik9FIFJ1bnRpbWVzIFN5c3RlbSBVc2VyIiwiZW1haWwiOiJydHN5c3VzckBjbi5pYm0uY29tIiwic3ViIjoicnRzeXN1c3JAY24uaWJtLmNvbSIsImFjY291bnQiOnsiYnNzIjoiOGQ2M2ZiMWNjNWU5OWU4NmRkNzIyOWRkZGZmYzA1YTUifSwiaWF0IjoxNTE2MTc0NjAzLCJleHAiOjE1MTYxNzgyMDMsImlzcyI6Imh0dHBzOi8vaWFtLmJsdWVtaXgubmV0L2lkZW50aXR5IiwiZ3JhbnRfdHlwZSI6InBhc3N3b3JkIiwic2NvcGUiOiJvcGVuaWQiLCJjbGllbnRfaWQiOiJieCJ9.gx-HQ1CSEwz5d4O1HXx4pusaYeEsqkQZgoBZ6esMBZG6wK6wQFPvC4D0Yvdi6CvKrVU-zV9PM_o3n5c-DFKjjTyTnRbQgrG0EPCRPmFW3bpepSb7eSw01S2YOLy5UTbz0cdM9hq-jafOu1S8pe9xeSMIMiA3-EFzCap5Z5CuoK9oIYJIFWseb1KsOyoiNOellbw1MaOmMzb4fsFz5Dr1Y8c1pNhoqp8M62E3y1yHe2jc6YepDab7Dqn2benK_e-MI3BlyWuBu4yo5mY2oCinJthr2E1YgbzWvcMy5a-ximnQIb4K6kscuUW_Yj_1GhDGJs4MP9u7M3-XdY1CNBGYeQ",
	"Bearer eyJraWQiOiIyMDE3MTAzMC0wMDowMDowMCIsImFsZyI6IlJTMjU2In0.eyJpYW1faWQiOiJJQk1pZC0yNzAwMDZWOEhNIiwiaWQiOiJJQk1pZC0yNzAwMDZWOEhNIiwicmVhbG1pZCI6IklCTWlkIiwiaWRlbnRpZmllciI6IjI3MDAwNlY4SE0iLCJnaXZlbl9uYW1lIjoiT0UgUnVudGltZXMiLCJmYW1pbHlfbmFtZSI6IlN5c3RlbSBVc2VyIiwibmFtZSI6Ik9FIFJ1bnRpbWVzIFN5c3RlbSBVc2VyIiwiZW1haWwiOiJydHN5c3VzckBjbi5pYm0uY29tIiwic3ViIjoicnRzeXN1c3JAY24uaWJtLmNvbSIsImFjY291bnQiOnsiYnNzIjoiOGQ2M2ZiMWNjNWU5OWU4NmRkNzIyOWRkZGZmYzA1YTUifSwiaWF0IjoxNTE2MTc0NjAzLCJleHAiOjE1MTYxNzgyMDMsImlzcyI6Imh0dHBzOi8vaWFtLmJsdWVtaXgubmV0L2lkZW50aXR5IiwiZ3JhbnRfdHlwZSI6InBhc3N3b3JkIiwic2NvcGUiOiJvcGVuaWQiLCJjbGllbnRfaWQiOiJieCJ9.gx-HQ1CSEwz5d4O1HXx4pusaYeEsqkQZgoBZ6esMBZG6wK6wQFPvC4D0Yvdi6CvKrVU-zV9PM_o3n5c-DFKjjTyTnRbQgrG0EPCRPmFW3bpepSb7eSw01S2YOLy5UTbz0cdM9hq-jafOu1S8pe9xeSMIMiA3-EFzCap5Z5CuoK9oIYJIFWseb1KsOyoiNOellbw1MaOmMzb4fsFz5Dr1Y8c1pNhoqp8M62E3y1yHe2jc6YepDab7Dqn2benK_e-MI3BlyWuBu4yo5mY2oCinJthr2E1YgbzWvcMy5a-ximnQIb4K6kscuUW_Yj_1GhDGJs4MP9u7M3-XdY1CNBGYeQ",
	//token from apikey
	"eyJraWQiOiIyMDE3MTAzMC0wMDowMDowMCIsImFsZyI6IlJTMjU2In0.eyJpYW1faWQiOiJJQk1pZC0yNzAwMDZWOEhNIiwiaWQiOiJJQk1pZC0yNzAwMDZWOEhNIiwicmVhbG1pZCI6IklCTWlkIiwiaWRlbnRpZmllciI6IjI3MDAwNlY4SE0iLCJnaXZlbl9uYW1lIjoiT0UgUnVudGltZXMiLCJmYW1pbHlfbmFtZSI6IlN5c3RlbSBVc2VyIiwibmFtZSI6Ik9FIFJ1bnRpbWVzIFN5c3RlbSBVc2VyIiwiZW1haWwiOiJydHN5c3VzckBjbi5pYm0uY29tIiwic3ViIjoicnRzeXN1c3JAY24uaWJtLmNvbSIsImFjY291bnQiOnsiYnNzIjoiOGQ2M2ZiMWNjNWU5OWU4NmRkNzIyOWRkZGZmYzA1YTUifSwiaWF0IjoxNTE2MTc1MDA3LCJleHAiOjE1MTYxNzg2MDcsImlzcyI6Imh0dHBzOi8vaWFtLmJsdWVtaXgubmV0L2lkZW50aXR5IiwiZ3JhbnRfdHlwZSI6InVybjppYm06cGFyYW1zOm9hdXRoOmdyYW50LXR5cGU6YXBpa2V5Iiwic2NvcGUiOiJvcGVuaWQiLCJjbGllbnRfaWQiOiJieCJ9.CuSOKifh4DvE__bjwDsn5BKmAHF2NaXznoiA1KG-2s2njbJs9nQdOJ3lkOnM77BqvLEpu2cwsmhi4Gsdy-MiJ6ACub0A5zyB-D95IXsGYa5tbFQBLbPpmFDAgAhLG5gXlVnU7nyIJN17Slm3pcWSNXEdWcsA1tgDkC9gQc_rpDhUfhnFeGA2LpvVMtRDolcOrbRuWN4NEbBOwdTbG5-6ijZ5Ag2z3lVmlQZ_6BLBCSVM8WlI8eIGICqCx0HYsmCiMlSqZ-4fkpg2DBYYYX_XsMQlamGynuPeoiBckJIyGEgsJD2egYN2bOUNLcn5htSCGxoJ4HJfXJ70_iCzmovb0w",
	"Bearer eyJraWQiOiIyMDE3MTAzMC0wMDowMDowMCIsImFsZyI6IlJTMjU2In0.eyJpYW1faWQiOiJJQk1pZC0yNzAwMDZWOEhNIiwiaWQiOiJJQk1pZC0yNzAwMDZWOEhNIiwicmVhbG1pZCI6IklCTWlkIiwiaWRlbnRpZmllciI6IjI3MDAwNlY4SE0iLCJnaXZlbl9uYW1lIjoiT0UgUnVudGltZXMiLCJmYW1pbHlfbmFtZSI6IlN5c3RlbSBVc2VyIiwibmFtZSI6Ik9FIFJ1bnRpbWVzIFN5c3RlbSBVc2VyIiwiZW1haWwiOiJydHN5c3VzckBjbi5pYm0uY29tIiwic3ViIjoicnRzeXN1c3JAY24uaWJtLmNvbSIsImFjY291bnQiOnsiYnNzIjoiOGQ2M2ZiMWNjNWU5OWU4NmRkNzIyOWRkZGZmYzA1YTUifSwiaWF0IjoxNTE2MTc1MDA3LCJleHAiOjE1MTYxNzg2MDcsImlzcyI6Imh0dHBzOi8vaWFtLmJsdWVtaXgubmV0L2lkZW50aXR5IiwiZ3JhbnRfdHlwZSI6InVybjppYm06cGFyYW1zOm9hdXRoOmdyYW50LXR5cGU6YXBpa2V5Iiwic2NvcGUiOiJvcGVuaWQiLCJjbGllbnRfaWQiOiJieCJ9.CuSOKifh4DvE__bjwDsn5BKmAHF2NaXznoiA1KG-2s2njbJs9nQdOJ3lkOnM77BqvLEpu2cwsmhi4Gsdy-MiJ6ACub0A5zyB-D95IXsGYa5tbFQBLbPpmFDAgAhLG5gXlVnU7nyIJN17Slm3pcWSNXEdWcsA1tgDkC9gQc_rpDhUfhnFeGA2LpvVMtRDolcOrbRuWN4NEbBOwdTbG5-6ijZ5Ag2z3lVmlQZ_6BLBCSVM8WlI8eIGICqCx0HYsmCiMlSqZ-4fkpg2DBYYYX_XsMQlamGynuPeoiBckJIyGEgsJD2egYN2bOUNLcn5htSCGxoJ4HJfXJ70_iCzmovb0w",
}

var TestIAMCRTokenData = []string{
	//token from compute resource
	"eyJraWQiOiIyMDE3MTAzMC0wMDowMDowMCIsImFsZyI6IlJTMjU2In0.ewoJImlhbV9pZCI6ICJpYW0tUHJvZmlsZS05NDQ5N2QwZC0yYWMzLTQxYmYtYTk5My1hNDlkMWIxNDYyN2MiLAoJImlkIjogIklCTWlkLXRlc3QiLAoJInJlYWxtaWQiOiAiaWFtIiwKCSJqdGkiOiAiMDRkMjBiMjUtZWUyZC00MDBmLTg2MjMtOGNkODA3MGI1NDY4IiwKCSJpZGVudGlmaWVyIjogIlByb2ZpbGUtOTQ0OTdkMGQtMmFjMy00MWJmLWE5OTMtYTQ5ZDFiMTQ2MjdjIiwKCSJuYW1lIjogIk15IFByb2ZpbGUiLAoJInN1YiI6ICJQcm9maWxlLTk0NDk3ZDBkLTJhYzMtNDFiZi1hOTkzLWE0OWQxYjE0NjI3YyIsCgkiYXV0aG4iOiB7CgkgICJzdWIiOiAiY3JuOnYxOnN0YWdpbmc6cHVibGljOmlhbS1pZGVudGl0eTo6YS8xOGUzMDIwNzQ5Y2U0NzQ0YjBiNDcyNDY2ZDYxZmRiNDo6Y29tcHV0ZXJlc291cmNlOkZha2UtQ29tcHV0ZS1SZXNvdXJjZSIsCgkgICJpYW1faWQiOiAiY3JuLWNybjp2MTpzdGFnaW5nOnB1YmxpYzppYW0taWRlbnRpdHk6OmEvMThlMzAyMDc0OWNlNDc0NGIwYjQ3MjQ2NmQ2MWZkYjQ6OmNvbXB1dGVyZXNvdXJjZTpGYWtlLUNvbXB1dGUtUmVzb3VyY2UiLAoJICAibmFtZSI6ICJteV9jb21wdXRlX3Jlc291cmNlIgoJfSwKCSJhY2NvdW50IjogewoJICAiYm91bmRhcnkiOiAiZ2xvYmFsIiwKCSAgInZhbGlkIjogdHJ1ZSwKCSAgImJzcyI6ICJmYWtlX2JzcyIKCX0sCgkiaWF0IjogMTYyOTkyOTQ2MywKCSJleHAiOiAxNjI5OTMzMDYzLAoJImlzcyI6ICJodHRwczovL2lhbS5jbG91ZC5pYm0uY29tL2lkZW50aXR5IiwKCSJncmFudF90eXBlIjogInVybjppYm06cGFyYW1zOm9hdXRoOmdyYW50LXR5cGU6Y3ItdG9rZW4iLAoJInNjb3BlIjogImlibSBvcGVuaWQiLAoJImNsaWVudF9pZCI6ICJieCIKICB9.CuSOKifh4DvE__bjwDsn5BKmAHF2NaXznoiA1KG-2s2njbJs9nQdOJ3lkOnM77BqvLEpu2cwsmhi4Gsdy-MiJ6ACub0A5zyB-D95IXsGYa5tbFQBLbPpmFDAgAhLG5gXlVnU7nyIJN17Slm3pcWSNXEdWcsA1tgDkC9gQc_rpDhUfhnFeGA2LpvVMtRDolcOrbRuWN4NEbBOwdTbG5-6ijZ5Ag2z3lVmlQZ_6BLBCSVM8WlI8eIGICqCx0HYsmCiMlSqZ-4fkpg2DBYYYX_XsMQlamGynuPeoiBckJIyGEgsJD2egYN2bOUNLcn5htSCGxoJ4HJfXJ70_iCzmovb0w",
	"Bearer eyJraWQiOiIyMDE3MTAzMC0wMDowMDowMCIsImFsZyI6IlJTMjU2In0.ewoJImlhbV9pZCI6ICJpYW0tUHJvZmlsZS05NDQ5N2QwZC0yYWMzLTQxYmYtYTk5My1hNDlkMWIxNDYyN2MiLAoJImlkIjogIklCTWlkLXRlc3QiLAoJInJlYWxtaWQiOiAiaWFtIiwKCSJqdGkiOiAiMDRkMjBiMjUtZWUyZC00MDBmLTg2MjMtOGNkODA3MGI1NDY4IiwKCSJpZGVudGlmaWVyIjogIlByb2ZpbGUtOTQ0OTdkMGQtMmFjMy00MWJmLWE5OTMtYTQ5ZDFiMTQ2MjdjIiwKCSJuYW1lIjogIk15IFByb2ZpbGUiLAoJInN1YiI6ICJQcm9maWxlLTk0NDk3ZDBkLTJhYzMtNDFiZi1hOTkzLWE0OWQxYjE0NjI3YyIsCgkiYXV0aG4iOiB7CgkgICJzdWIiOiAiY3JuOnYxOnN0YWdpbmc6cHVibGljOmlhbS1pZGVudGl0eTo6YS8xOGUzMDIwNzQ5Y2U0NzQ0YjBiNDcyNDY2ZDYxZmRiNDo6Y29tcHV0ZXJlc291cmNlOkZha2UtQ29tcHV0ZS1SZXNvdXJjZSIsCgkgICJpYW1faWQiOiAiY3JuLWNybjp2MTpzdGFnaW5nOnB1YmxpYzppYW0taWRlbnRpdHk6OmEvMThlMzAyMDc0OWNlNDc0NGIwYjQ3MjQ2NmQ2MWZkYjQ6OmNvbXB1dGVyZXNvdXJjZTpGYWtlLUNvbXB1dGUtUmVzb3VyY2UiLAoJICAibmFtZSI6ICJteV9jb21wdXRlX3Jlc291cmNlIgoJfSwKCSJhY2NvdW50IjogewoJICAiYm91bmRhcnkiOiAiZ2xvYmFsIiwKCSAgInZhbGlkIjogdHJ1ZSwKCSAgImJzcyI6ICJmYWtlX2JzcyIKCX0sCgkiaWF0IjogMTYyOTkyOTQ2MywKCSJleHAiOiAxNjI5OTMzMDYzLAoJImlzcyI6ICJodHRwczovL2lhbS5jbG91ZC5pYm0uY29tL2lkZW50aXR5IiwKCSJncmFudF90eXBlIjogInVybjppYm06cGFyYW1zOm9hdXRoOmdyYW50LXR5cGU6Y3ItdG9rZW4iLAoJInNjb3BlIjogImlibSBvcGVuaWQiLAoJImNsaWVudF9pZCI6ICJieCIKICB9.CuSOKifh4DvE__bjwDsn5BKmAHF2NaXznoiA1KG-2s2njbJs9nQdOJ3lkOnM77BqvLEpu2cwsmhi4Gsdy-MiJ6ACub0A5zyB-D95IXsGYa5tbFQBLbPpmFDAgAhLG5gXlVnU7nyIJN17Slm3pcWSNXEdWcsA1tgDkC9gQc_rpDhUfhnFeGA2LpvVMtRDolcOrbRuWN4NEbBOwdTbG5-6ijZ5Ag2z3lVmlQZ_6BLBCSVM8WlI8eIGICqCx0HYsmCiMlSqZ-4fkpg2DBYYYX_XsMQlamGynuPeoiBckJIyGEgsJD2egYN2bOUNLcn5htSCGxoJ4HJfXJ70_iCzmovb0w",
}

var TestUAATokenData = []string{
	"Bearer eyJhbGciOiJIUzI1NiIsImtpZCI6ImtleS0xIiwidHlwIjoiSldUIn0.eyJqdGkiOiJhZGFkNWYwYzQ0ZDI0ZDlkYmVhN2YyNGIzMDNmOWNhNyIsInN1YiI6IjY3ODdiMzM2LTAwNzUtNGYwYy1hZmZiLWUyOWZjMmVhZWI4OCIsInNjb3BlIjpbIm9wZW5pZCIsInVhYS51c2VyIiwiY2xvdWRfY29udHJvbGxlci5yZWFkIiwicGFzc3dvcmQud3JpdGUiLCJjbG91ZF9jb250cm9sbGVyLndyaXRlIl0sImNsaWVudF9pZCI6ImNmIiwiY2lkIjoiY2YiLCJhenAiOiJjZiIsImdyYW50X3R5cGUiOiJwYXNzd29yZCIsInVzZXJfaWQiOiI2Nzg3YjMzNi0wMDc1LTRmMGMtYWZmYi1lMjlmYzJlYWViODgiLCJvcmlnaW4iOiJ1YWEiLCJ1c2VyX25hbWUiOiJ3YW5nanVubEBjbi5pYm0uY29tIiwiZW1haWwiOiJ3YW5nanVubEBjbi5pYm0uY29tIiwiYXV0aF90aW1lIjoxNTE2MTczMjgxLCJpYXQiOjE1MTYxNzMyODEsImV4cCI6MTUxNjE3Njg4MCwiaXNzIjoiaHR0cHM6Ly91YWEubmcuYmx1ZW1peC5uZXQvb2F1dGgvdG9rZW4iLCJ6aWQiOiJ1YWEiLCJhdWQiOlsiY2xvdWRfY29udHJvbGxlciIsInBhc3N3b3JkIiwiY2YiLCJ1YWEiLCJvcGVuaWQiXX0.K-9HdMCrNdln81ewX_TQLR63F4wChz035G5KtMq9wkk",
}

func TestNewIAMTokenInfo(t *testing.T) {
	for _, token := range TestIAMTokenData {
		tokenInfo := NewIAMTokenInfo(token)
		assert.Equal(t, tokenInfo.UserEmail, "rtsysusr@cn.ibm.com")
		assert.Equal(t, tokenInfo.IAMID, "IBMid-270006V8HM")
		assert.Equal(t, tokenInfo.Accounts.AccountID, "8d63fb1cc5e99e86dd7229dddffc05a5")
	}
}
func TestNewIAMCRTokenInfo(t *testing.T) {
	for _, token := range TestIAMCRTokenData {
		tokenInfo := NewIAMTokenInfo(token)
		assert.Equal(t, tokenInfo.IAMID, "iam-Profile-94497d0d-2ac3-41bf-a993-a49d1b14627c")
		assert.Equal(t, tokenInfo.Fullname, "My Profile")
		assert.Equal(t, tokenInfo.Subject, "Profile-94497d0d-2ac3-41bf-a993-a49d1b14627c")
		assert.Equal(t, tokenInfo.Authn.Subject, "crn:v1:staging:public:iam-identity::a/18e3020749ce4744b0b472466d61fdb4::computeresource:Fake-Compute-Resource")
		assert.Equal(t, tokenInfo.Authn.IAMID, "crn-crn:v1:staging:public:iam-identity::a/18e3020749ce4744b0b472466d61fdb4::computeresource:Fake-Compute-Resource")
		assert.Equal(t, tokenInfo.Authn.Name, "my_compute_resource")
	}
}
func TestIATandEXP(t *testing.T) {
	tokenInfo := NewIAMTokenInfo(TestIAMTokenData[1])
	assert.Equal(t, tokenInfo.IssueAt.Unix(), int64(1516174603))
	assert.Equal(t, tokenInfo.Expiry.Unix(), int64(1516178203))
}

func TestUAATokenInfo(t *testing.T) {
	for _, token := range TestUAATokenData {
		tokenInfo := NewUAATokenInfo(token)
		assert.Equal(t, tokenInfo.Username, "wangjunl@cn.ibm.com")
		assert.Equal(t, tokenInfo.UserGUID, "6787b336-0075-4f0c-affb-e29fc2eaeb88")
	}
}
