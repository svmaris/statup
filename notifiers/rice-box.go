// Statup
// Copyright (C) 2018.  Hunter Long and the project contributors
// Written by Hunter Long <info@socialeck.com> and the project contributors
//
// https://github.com/hunterlong/statup
//
// The licenses for most software and other practical works are designed
// to take away your freedom to share and change the works.  By contrast,
// the GNU General Public License is intended to guarantee your freedom to
// share and change all versions of a program--to make sure it remains free
// software for all its users.
//
// You should have received a copy of the GNU General Public License
// along with this program.  If not, see <http://www.gnu.org/licenses/>.

package notifiers

import (
	"github.com/GeertJohan/go.rice/embedded"
	"time"
)

func init() {

	// define files
	file2 := &embedded.EmbeddedFile{
		Filename:    "failure.html",
		FileModTime: time.Unix(1531720141, 0),
		Content:     string("<!DOCTYPE html PUBLIC \"-//W3C//DTD XHTML 1.0 Transitional//EN\" \"http://www.w3.org/TR/xhtml1/DTD/xhtml1-transitional.dtd\">\n<html xmlns=\"http://www.w3.org/1999/xhtml\" xmlns=\"http://www.w3.org/1999/xhtml\">\n<head>\n    <meta name=\"viewport\" content=\"width=device-width, initial-scale=1.0\" />\n    <meta http-equiv=\"Content-Type\" content=\"text/html; charset=UTF-8\" />\n    <title>Sample Email</title>\n\n\n</head>\n<body style=\"-webkit-text-size-adjust: none; box-sizing: border-box; color: #74787E; font-family: Arial, 'Helvetica Neue', Helvetica, sans-serif; height: 100%; line-height: 1.4; margin: 0; width: 100% !important;\" bgcolor=\"#F2F4F6\"><style type=\"text/css\">\n    body {\n        width: 100% !important; height: 100%; margin: 0; line-height: 1.4; background-color: #F2F4F6; color: #74787E; -webkit-text-size-adjust: none;\n    }\n    @media only screen and (max-width: 600px) {\n        .email-body_inner {\n            width: 100% !important;\n        }\n        .email-footer {\n            width: 100% !important;\n        }\n    }\n    @media only screen and (max-width: 500px) {\n        .button {\n            width: 100% !important;\n        }\n    }\n</style>\n<table class=\"email-wrapper\" width=\"100%\" cellpadding=\"0\" cellspacing=\"0\" style=\"box-sizing: border-box; font-family: Arial, 'Helvetica Neue', Helvetica, sans-serif; margin: 0; padding: 0; width: 100%;\" bgcolor=\"#F2F4F6\">\n    <tr>\n        <td align=\"center\" style=\"box-sizing: border-box; font-family: Arial, 'Helvetica Neue', Helvetica, sans-serif; word-break: break-word;\">\n            <table class=\"email-content\" width=\"100%\" cellpadding=\"0\" cellspacing=\"0\" style=\"box-sizing: border-box; font-family: Arial, 'Helvetica Neue', Helvetica, sans-serif; margin: 0; padding: 0; width: 100%;\">\n\n                <tr>\n                    <td class=\"email-body\" width=\"100%\" cellpadding=\"0\" cellspacing=\"0\" style=\"-premailer-cellpadding: 0; -premailer-cellspacing: 0; border-bottom-color: #EDEFF2; border-bottom-style: solid; border-bottom-width: 1px; border-top-color: #EDEFF2; border-top-style: solid; border-top-width: 1px; box-sizing: border-box; font-family: Arial, 'Helvetica Neue', Helvetica, sans-serif; margin: 0; padding: 0; width: 100%; word-break: break-word;\" bgcolor=\"#FFFFFF\">\n                        <table class=\"email-body_inner\" align=\"center\" width=\"570\" cellpadding=\"0\" cellspacing=\"0\" style=\"box-sizing: border-box; font-family: Arial, 'Helvetica Neue', Helvetica, sans-serif; margin: 0 auto; padding: 0; width: 570px;\" bgcolor=\"#FFFFFF\">\n\n                            <tr>\n                                <td class=\"content-cell\" style=\"box-sizing: border-box; font-family: Arial, 'Helvetica Neue', Helvetica, sans-serif; padding: 35px; word-break: break-word;\">\n                                    <h1 style=\"box-sizing: border-box; color: #2F3133; font-family: Arial, 'Helvetica Neue', Helvetica, sans-serif; font-size: 19px; font-weight: bold; margin-top: 0;\" align=\"left\">{{ .Service.Name }} is Offline!</h1>\n                                    <p style=\"box-sizing: border-box; color: #74787E; font-family: Arial, 'Helvetica Neue', Helvetica, sans-serif; font-size: 16px; line-height: 1.5em; margin-top: 0;\" align=\"left\">\n                                        Your Statup service '<a target=\"_blank\" href=\"{{.Service.Domain}}\">{{.Service.Name}}</a>' has been triggered with a HTTP status code of '{{.Service.LastStatusCode}}' and is currently offline based on your requirements. This failure was created on {{.Service.CreatedAt}}.\n                                    </p>\n\n                                {{if .Service.LastResponse }}\n                                    <h1 style=\"box-sizing: border-box; color: #2F3133; font-family: Arial, 'Helvetica Neue', Helvetica, sans-serif; font-size: 19px; font-weight: bold; margin-top: 0;\" align=\"left\">Last Response</h1>\n                                    <p style=\"box-sizing: border-box; color: #74787E; font-family: Arial, 'Helvetica Neue', Helvetica, sans-serif; font-size: 16px; line-height: 1.5em; margin-top: 0;\" align=\"left\">\n                                        {{ .Service.LastResponse }}\n                                    </p>\n                                {{end}}\n\n                                    <table class=\"body-sub\" style=\"border-top-color: #EDEFF2; border-top-style: solid; border-top-width: 1px; box-sizing: border-box; font-family: Arial, 'Helvetica Neue', Helvetica, sans-serif; margin-top: 25px; padding-top: 25px;\">\n                                        <td style=\"box-sizing: border-box; font-family: Arial, 'Helvetica Neue', Helvetica, sans-serif; word-break: break-word;\">\n                                            <a href=\"/service/{{.Service.Id}}\" class=\"button button--blue\" target=\"_blank\" style=\"-webkit-text-size-adjust: none; background: #3869D4; border-color: #3869d4; border-radius: 3px; border-style: solid; border-width: 10px 18px; box-shadow: 0 2px 3px rgba(0, 0, 0, 0.16); box-sizing: border-box; color: #FFF; display: inline-block; font-family: Arial, 'Helvetica Neue', Helvetica, sans-serif; text-decoration: none;\">View Service</a>\n                                        </td>\n                                        <td style=\"box-sizing: border-box; font-family: Arial, 'Helvetica Neue', Helvetica, sans-serif; word-break: break-word;\">\n                                            <a href=\"/dashboard\" class=\"button button--blue\" target=\"_blank\" style=\"-webkit-text-size-adjust: none; background: #3869D4; border-color: #3869d4; border-radius: 3px; border-style: solid; border-width: 10px 18px; box-shadow: 0 2px 3px rgba(0, 0, 0, 0.16); box-sizing: border-box; color: #FFF; display: inline-block; font-family: Arial, 'Helvetica Neue', Helvetica, sans-serif; text-decoration: none;\">Statup Dashboard</a>\n                                        </td>\n                                    </table>\n                                </td>\n                            </tr>\n                        </table>\n                    </td>\n                </tr>\n            </table>\n        </td>\n    </tr>\n</table>\n</body>\n</html>"),
	}
	file3 := &embedded.EmbeddedFile{
		Filename:    "message.html",
		FileModTime: time.Unix(1530546686, 0),
		Content:     string("<!DOCTYPE html PUBLIC \"-//W3C//DTD XHTML 1.0 Transitional//EN\" \"http://www.w3.org/TR/xhtml1/DTD/xhtml1-transitional.dtd\">\n<html xmlns=\"http://www.w3.org/1999/xhtml\" xmlns=\"http://www.w3.org/1999/xhtml\">\n<head>\n    <meta name=\"viewport\" content=\"width=device-width, initial-scale=1.0\" />\n    <meta http-equiv=\"Content-Type\" content=\"text/html; charset=UTF-8\" />\n    <title>Sample Email</title>\n\n\n</head>\n<body style=\"-webkit-text-size-adjust: none; box-sizing: border-box; color: #74787E; font-family: Arial, 'Helvetica Neue', Helvetica, sans-serif; height: 100%; line-height: 1.4; margin: 0; width: 100% !important;\" bgcolor=\"#F2F4F6\"><style type=\"text/css\">\n    body {\n        width: 100% !important; height: 100%; margin: 0; line-height: 1.4; background-color: #F2F4F6; color: #74787E; -webkit-text-size-adjust: none;\n    }\n    @media only screen and (max-width: 600px) {\n        .email-body_inner {\n            width: 100% !important;\n        }\n        .email-footer {\n            width: 100% !important;\n        }\n    }\n    @media only screen and (max-width: 500px) {\n        .button {\n            width: 100% !important;\n        }\n    }\n</style>\n<table class=\"email-wrapper\" width=\"100%\" cellpadding=\"0\" cellspacing=\"0\" style=\"box-sizing: border-box; font-family: Arial, 'Helvetica Neue', Helvetica, sans-serif; margin: 0; padding: 0; width: 100%;\" bgcolor=\"#F2F4F6\">\n    <tr>\n        <td align=\"center\" style=\"box-sizing: border-box; font-family: Arial, 'Helvetica Neue', Helvetica, sans-serif; word-break: break-word;\">\n            <table class=\"email-content\" width=\"100%\" cellpadding=\"0\" cellspacing=\"0\" style=\"box-sizing: border-box; font-family: Arial, 'Helvetica Neue', Helvetica, sans-serif; margin: 0; padding: 0; width: 100%;\">\n\n                <tr>\n                    <td class=\"email-body\" width=\"100%\" cellpadding=\"0\" cellspacing=\"0\" style=\"-premailer-cellpadding: 0; -premailer-cellspacing: 0; border-bottom-color: #EDEFF2; border-bottom-style: solid; border-bottom-width: 1px; border-top-color: #EDEFF2; border-top-style: solid; border-top-width: 1px; box-sizing: border-box; font-family: Arial, 'Helvetica Neue', Helvetica, sans-serif; margin: 0; padding: 0; width: 100%; word-break: break-word;\" bgcolor=\"#FFFFFF\">\n                        <table class=\"email-body_inner\" align=\"center\" width=\"570\" cellpadding=\"0\" cellspacing=\"0\" style=\"box-sizing: border-box; font-family: Arial, 'Helvetica Neue', Helvetica, sans-serif; margin: 0 auto; padding: 0; width: 570px;\" bgcolor=\"#FFFFFF\">\n\n                            <tr>\n                                <td class=\"content-cell\" style=\"box-sizing: border-box; font-family: Arial, 'Helvetica Neue', Helvetica, sans-serif; padding: 35px; word-break: break-word;\">\n                                    <h1 style=\"box-sizing: border-box; color: #2F3133; font-family: Arial, 'Helvetica Neue', Helvetica, sans-serif; font-size: 19px; font-weight: bold; margin-top: 0;\" align=\"left\">Looks Like Emails Work!</h1>\n                                    <p style=\"box-sizing: border-box; color: #74787E; font-family: Arial, 'Helvetica Neue', Helvetica, sans-serif; font-size: 16px; line-height: 1.5em; margin-top: 0;\" align=\"left\">\n                                        Since you got this email, it confirms that your Statup Status Page email system is working correctly.\n                                    </p>\n                                    </p>\n                                    <p style=\"box-sizing: border-box; color: #74787E; font-family: Arial, 'Helvetica Neue', Helvetica, sans-serif; font-size: 16px; line-height: 1.5em; margin-top: 0;\" align=\"left\">\n                                        Enjoy using Statup!\n                                        <br />Statup.io Team</p>\n\n                                    <table class=\"body-sub\" style=\"border-top-color: #EDEFF2; border-top-style: solid; border-top-width: 1px; box-sizing: border-box; font-family: Arial, 'Helvetica Neue', Helvetica, sans-serif; margin-top: 25px; padding-top: 25px;\">\n\n                                    </table>\n                                </td>\n                            </tr>\n                        </table>\n                    </td>\n                </tr>\n            </table>\n        </td>\n    </tr>\n</table>\n</body>\n</html>"),
	}

	// define dirs
	dir1 := &embedded.EmbeddedDir{
		Filename:   "",
		DirModTime: time.Unix(1531720141, 0),
		ChildFiles: []*embedded.EmbeddedFile{
			file2, // "failure.html"
			file3, // "message.html"

		},
	}

	// link ChildDirs
	dir1.ChildDirs = []*embedded.EmbeddedDir{}

	// register embeddedBox
	embedded.RegisterEmbeddedBox(`emails`, &embedded.EmbeddedBox{
		Name: `emails`,
		Time: time.Unix(1531720141, 0),
		Dirs: map[string]*embedded.EmbeddedDir{
			"": dir1,
		},
		Files: map[string]*embedded.EmbeddedFile{
			"failure.html": file2,
			"message.html": file3,
		},
	})
}
