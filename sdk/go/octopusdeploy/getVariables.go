// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package octopusdeploy

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Provides information about existing variables.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-octopusdeploy/sdk/go/octopusdeploy"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := octopusdeploy.GetVariables(ctx, nil, nil)
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
func GetVariables(ctx *pulumi.Context, args *GetVariablesArgs, opts ...pulumi.InvokeOption) (*GetVariablesResult, error) {
	var rv GetVariablesResult
	err := ctx.Invoke("octopusdeploy:index/getVariables:getVariables", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getVariables.
type GetVariablesArgs struct {
	// A filter to search by a list of IDs.
	Ids []string `pulumi:"ids"`
	// A list of variables that match the filter(s).
	Variables []GetVariablesVariable `pulumi:"variables"`
}

// A collection of values returned by getVariables.
type GetVariablesResult struct {
	// An auto-generated identifier that includes the timestamp when this data source was last modified.
	Id string `pulumi:"id"`
	// A filter to search by a list of IDs.
	Ids []string `pulumi:"ids"`
	// A list of variables that match the filter(s).
	Variables []GetVariablesVariable `pulumi:"variables"`
}

func GetVariablesOutput(ctx *pulumi.Context, args GetVariablesOutputArgs, opts ...pulumi.InvokeOption) GetVariablesResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetVariablesResult, error) {
			args := v.(GetVariablesArgs)
			r, err := GetVariables(ctx, &args, opts...)
			var s GetVariablesResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(GetVariablesResultOutput)
}

// A collection of arguments for invoking getVariables.
type GetVariablesOutputArgs struct {
	// A filter to search by a list of IDs.
	Ids pulumi.StringArrayInput `pulumi:"ids"`
	// A list of variables that match the filter(s).
	Variables GetVariablesVariableArrayInput `pulumi:"variables"`
}

func (GetVariablesOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetVariablesArgs)(nil)).Elem()
}

// A collection of values returned by getVariables.
type GetVariablesResultOutput struct{ *pulumi.OutputState }

func (GetVariablesResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetVariablesResult)(nil)).Elem()
}

func (o GetVariablesResultOutput) ToGetVariablesResultOutput() GetVariablesResultOutput {
	return o
}

func (o GetVariablesResultOutput) ToGetVariablesResultOutputWithContext(ctx context.Context) GetVariablesResultOutput {
	return o
}

// An auto-generated identifier that includes the timestamp when this data source was last modified.
func (o GetVariablesResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetVariablesResult) string { return v.Id }).(pulumi.StringOutput)
}

// A filter to search by a list of IDs.
func (o GetVariablesResultOutput) Ids() pulumi.StringArrayOutput {
	return o.ApplyT(func(v GetVariablesResult) []string { return v.Ids }).(pulumi.StringArrayOutput)
}

// A list of variables that match the filter(s).
func (o GetVariablesResultOutput) Variables() GetVariablesVariableArrayOutput {
	return o.ApplyT(func(v GetVariablesResult) []GetVariablesVariable { return v.Variables }).(GetVariablesVariableArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(GetVariablesResultOutput{})
}