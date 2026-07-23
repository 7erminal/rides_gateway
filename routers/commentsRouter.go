package routers

import (
	beego "github.com/beego/beego/v2/server/web"
	"github.com/beego/beego/v2/server/web/context/param"
)

func init() {

    beego.GlobalControllerRouter["AMC_gateway/controllers:AuthenticationController"] = append(beego.GlobalControllerRouter["AMC_gateway/controllers:AuthenticationController"],
        beego.ControllerComments{
            Method: "ChangePassword",
            Router: `/change-password`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["AMC_gateway/controllers:AuthenticationController"] = append(beego.GlobalControllerRouter["AMC_gateway/controllers:AuthenticationController"],
        beego.ControllerComments{
            Method: "Register",
            Router: `/register`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["AMC_gateway/controllers:AuthenticationController"] = append(beego.GlobalControllerRouter["AMC_gateway/controllers:AuthenticationController"],
        beego.ControllerComments{
            Method: "SendActivationCode",
            Router: `/send-activation-code`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["AMC_gateway/controllers:AuthenticationController"] = append(beego.GlobalControllerRouter["AMC_gateway/controllers:AuthenticationController"],
        beego.ControllerComments{
            Method: "SignIn",
            Router: `/sign-in`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["AMC_gateway/controllers:AuthenticationController"] = append(beego.GlobalControllerRouter["AMC_gateway/controllers:AuthenticationController"],
        beego.ControllerComments{
            Method: "VerifyActivationCode",
            Router: `/verify-activation-code`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["AMC_gateway/controllers:CustomermanagementController"] = append(beego.GlobalControllerRouter["AMC_gateway/controllers:CustomermanagementController"],
        beego.ControllerComments{
            Method: "Post",
            Router: `/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["AMC_gateway/controllers:CustomermanagementController"] = append(beego.GlobalControllerRouter["AMC_gateway/controllers:CustomermanagementController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["AMC_gateway/controllers:CustomermanagementController"] = append(beego.GlobalControllerRouter["AMC_gateway/controllers:CustomermanagementController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: `/:id`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["AMC_gateway/controllers:CustomermanagementController"] = append(beego.GlobalControllerRouter["AMC_gateway/controllers:CustomermanagementController"],
        beego.ControllerComments{
            Method: "Put",
            Router: `/:id`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["AMC_gateway/controllers:CustomermanagementController"] = append(beego.GlobalControllerRouter["AMC_gateway/controllers:CustomermanagementController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: `/:id`,
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["AMC_gateway/controllers:CustomermanagementController"] = append(beego.GlobalControllerRouter["AMC_gateway/controllers:CustomermanagementController"],
        beego.ControllerComments{
            Method: "UpdateCustomerImage",
            Router: `/upload-customer-image`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["AMC_gateway/controllers:ItemsController"] = append(beego.GlobalControllerRouter["AMC_gateway/controllers:ItemsController"],
        beego.ControllerComments{
            Method: "AddCategory",
            Router: `/add-product-type`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["AMC_gateway/controllers:ItemsController"] = append(beego.GlobalControllerRouter["AMC_gateway/controllers:ItemsController"],
        beego.ControllerComments{
            Method: "AddRentalsItem",
            Router: `/add-rental-product`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["AMC_gateway/controllers:ItemsController"] = append(beego.GlobalControllerRouter["AMC_gateway/controllers:ItemsController"],
        beego.ControllerComments{
            Method: "AddSalesItem",
            Router: `/add-sales-product`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["AMC_gateway/controllers:ItemsController"] = append(beego.GlobalControllerRouter["AMC_gateway/controllers:ItemsController"],
        beego.ControllerComments{
            Method: "GetCategories",
            Router: `/get-product-types`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["AMC_gateway/controllers:ItemsController"] = append(beego.GlobalControllerRouter["AMC_gateway/controllers:ItemsController"],
        beego.ControllerComments{
            Method: "GetProduct",
            Router: `/get-product/:id`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["AMC_gateway/controllers:ItemsController"] = append(beego.GlobalControllerRouter["AMC_gateway/controllers:ItemsController"],
        beego.ControllerComments{
            Method: "GetItems",
            Router: `/get-products`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["AMC_gateway/controllers:ItemsController"] = append(beego.GlobalControllerRouter["AMC_gateway/controllers:ItemsController"],
        beego.ControllerComments{
            Method: "UpdateItem",
            Router: `/update-product/:id`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["AMC_gateway/controllers:ItemsController"] = append(beego.GlobalControllerRouter["AMC_gateway/controllers:ItemsController"],
        beego.ControllerComments{
            Method: "UpdateItemImage",
            Router: `/upload-product-image`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["AMC_gateway/controllers:PaymentController"] = append(beego.GlobalControllerRouter["AMC_gateway/controllers:PaymentController"],
        beego.ControllerComments{
            Method: "GetPaymentMethods",
            Router: `/payment-methods`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["AMC_gateway/controllers:PaymentController"] = append(beego.GlobalControllerRouter["AMC_gateway/controllers:PaymentController"],
        beego.ControllerComments{
            Method: "UploadPaymentProof",
            Router: `/upload-payment-proof`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["AMC_gateway/controllers:RegistrationController"] = append(beego.GlobalControllerRouter["AMC_gateway/controllers:RegistrationController"],
        beego.ControllerComments{
            Method: "RegisterAlt",
            Router: `/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["AMC_gateway/controllers:StatsController"] = append(beego.GlobalControllerRouter["AMC_gateway/controllers:StatsController"],
        beego.ControllerComments{
            Method: "GetGeneralStats",
            Router: `/get-stats`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["AMC_gateway/controllers:SystemController"] = append(beego.GlobalControllerRouter["AMC_gateway/controllers:SystemController"],
        beego.ControllerComments{
            Method: "AddBranch",
            Router: `/add-branch`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["AMC_gateway/controllers:SystemController"] = append(beego.GlobalControllerRouter["AMC_gateway/controllers:SystemController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: `/delete-branch/:id`,
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["AMC_gateway/controllers:SystemController"] = append(beego.GlobalControllerRouter["AMC_gateway/controllers:SystemController"],
        beego.ControllerComments{
            Method: "GetOneBranch",
            Router: `/get-branch/:id`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["AMC_gateway/controllers:SystemController"] = append(beego.GlobalControllerRouter["AMC_gateway/controllers:SystemController"],
        beego.ControllerComments{
            Method: "GetAllBranches",
            Router: `/get-branches`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["AMC_gateway/controllers:SystemController"] = append(beego.GlobalControllerRouter["AMC_gateway/controllers:SystemController"],
        beego.ControllerComments{
            Method: "GetAllCountries",
            Router: `/get-countries`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["AMC_gateway/controllers:SystemController"] = append(beego.GlobalControllerRouter["AMC_gateway/controllers:SystemController"],
        beego.ControllerComments{
            Method: "GetIdTypes",
            Router: `/get-id-types`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["AMC_gateway/controllers:SystemController"] = append(beego.GlobalControllerRouter["AMC_gateway/controllers:SystemController"],
        beego.ControllerComments{
            Method: "UpdateBranch",
            Router: `/update-branch/:id`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["AMC_gateway/controllers:TransactionsController"] = append(beego.GlobalControllerRouter["AMC_gateway/controllers:TransactionsController"],
        beego.ControllerComments{
            Method: "PlaceRentalRequest",
            Router: `/place-rental-request`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["AMC_gateway/controllers:TransactionsController"] = append(beego.GlobalControllerRouter["AMC_gateway/controllers:TransactionsController"],
        beego.ControllerComments{
            Method: "PlaceSalesRequest",
            Router: `/place-sales-request`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["AMC_gateway/controllers:UserManagementController"] = append(beego.GlobalControllerRouter["AMC_gateway/controllers:UserManagementController"],
        beego.ControllerComments{
            Method: "GetBranchManagers",
            Router: `/get-branch-managers`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["AMC_gateway/controllers:UserManagementController"] = append(beego.GlobalControllerRouter["AMC_gateway/controllers:UserManagementController"],
        beego.ControllerComments{
            Method: "GetUserInvites",
            Router: `/get-invites`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["AMC_gateway/controllers:UserManagementController"] = append(beego.GlobalControllerRouter["AMC_gateway/controllers:UserManagementController"],
        beego.ControllerComments{
            Method: "GetRoles",
            Router: `/get-roles`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["AMC_gateway/controllers:UserManagementController"] = append(beego.GlobalControllerRouter["AMC_gateway/controllers:UserManagementController"],
        beego.ControllerComments{
            Method: "GetUserInvite",
            Router: `/get-user-invite/:token`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["AMC_gateway/controllers:UserManagementController"] = append(beego.GlobalControllerRouter["AMC_gateway/controllers:UserManagementController"],
        beego.ControllerComments{
            Method: "GetUser",
            Router: `/get-user-session`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["AMC_gateway/controllers:UserManagementController"] = append(beego.GlobalControllerRouter["AMC_gateway/controllers:UserManagementController"],
        beego.ControllerComments{
            Method: "GetUserWithId",
            Router: `/get-user-with-id/:id`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["AMC_gateway/controllers:UserManagementController"] = append(beego.GlobalControllerRouter["AMC_gateway/controllers:UserManagementController"],
        beego.ControllerComments{
            Method: "GetUsers",
            Router: `/get-users`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["AMC_gateway/controllers:UserManagementController"] = append(beego.GlobalControllerRouter["AMC_gateway/controllers:UserManagementController"],
        beego.ControllerComments{
            Method: "GetUsersUnderBranch",
            Router: `/get-users-under-branch/:branch_id`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["AMC_gateway/controllers:UserManagementController"] = append(beego.GlobalControllerRouter["AMC_gateway/controllers:UserManagementController"],
        beego.ControllerComments{
            Method: "InviteUserReg",
            Router: `/invite-user`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["AMC_gateway/controllers:UserManagementController"] = append(beego.GlobalControllerRouter["AMC_gateway/controllers:UserManagementController"],
        beego.ControllerComments{
            Method: "LogOut",
            Router: `/log-out`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["AMC_gateway/controllers:UserManagementController"] = append(beego.GlobalControllerRouter["AMC_gateway/controllers:UserManagementController"],
        beego.ControllerComments{
            Method: "UpdateInviteToken",
            Router: `/revoke-invite/:id`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["AMC_gateway/controllers:UserManagementController"] = append(beego.GlobalControllerRouter["AMC_gateway/controllers:UserManagementController"],
        beego.ControllerComments{
            Method: "UpdateUserBranch",
            Router: `/update-user-branch/:userid`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["AMC_gateway/controllers:UserManagementController"] = append(beego.GlobalControllerRouter["AMC_gateway/controllers:UserManagementController"],
        beego.ControllerComments{
            Method: "UpdateUserImage",
            Router: `/update-user-image`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["AMC_gateway/controllers:UserManagementController"] = append(beego.GlobalControllerRouter["AMC_gateway/controllers:UserManagementController"],
        beego.ControllerComments{
            Method: "UpdateUserRole",
            Router: `/update-user-role/:userid`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["AMC_gateway/controllers:UserManagementController"] = append(beego.GlobalControllerRouter["AMC_gateway/controllers:UserManagementController"],
        beego.ControllerComments{
            Method: "UpdateUser",
            Router: `/update-user/:id`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["AMC_gateway/controllers:UserManagementController"] = append(beego.GlobalControllerRouter["AMC_gateway/controllers:UserManagementController"],
        beego.ControllerComments{
            Method: "VerifyInvite",
            Router: `/verify-invite`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

}
