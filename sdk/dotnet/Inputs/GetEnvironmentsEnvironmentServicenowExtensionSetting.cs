// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Octopusdeploy.Inputs
{

    public sealed class GetEnvironmentsEnvironmentServicenowExtensionSettingArgs : global::Pulumi.InvokeArgs
    {
        [Input("isEnabled", required: true)]
        public bool IsEnabled { get; set; }

        public GetEnvironmentsEnvironmentServicenowExtensionSettingArgs()
        {
        }
        public static new GetEnvironmentsEnvironmentServicenowExtensionSettingArgs Empty => new GetEnvironmentsEnvironmentServicenowExtensionSettingArgs();
    }
}
