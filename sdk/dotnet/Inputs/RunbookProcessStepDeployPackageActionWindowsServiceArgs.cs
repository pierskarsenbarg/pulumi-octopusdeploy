// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace PiersKarsenbarg.Octopusdeploy.Inputs
{

    public sealed class RunbookProcessStepDeployPackageActionWindowsServiceArgs : global::Pulumi.ResourceArgs
    {
        [Input("arguments")]
        public Input<string>? Arguments { get; set; }

        [Input("createOrUpdateService")]
        public Input<bool>? CreateOrUpdateService { get; set; }

        [Input("customAccountName")]
        public Input<string>? CustomAccountName { get; set; }

        [Input("customAccountPassword")]
        private Input<string>? _customAccountPassword;
        public Input<string>? CustomAccountPassword
        {
            get => _customAccountPassword;
            set
            {
                var emptySecret = Output.CreateSecret(0);
                _customAccountPassword = Output.Tuple<Input<string>?, int>(value, emptySecret).Apply(t => t.Item1);
            }
        }

        [Input("dependencies")]
        public Input<string>? Dependencies { get; set; }

        [Input("description")]
        public Input<string>? Description { get; set; }

        [Input("displayName")]
        public Input<string>? DisplayName { get; set; }

        [Input("executablePath", required: true)]
        public Input<string> ExecutablePath { get; set; } = null!;

        [Input("serviceAccount")]
        public Input<string>? ServiceAccount { get; set; }

        [Input("serviceName", required: true)]
        public Input<string> ServiceName { get; set; } = null!;

        [Input("startMode")]
        public Input<string>? StartMode { get; set; }

        public RunbookProcessStepDeployPackageActionWindowsServiceArgs()
        {
        }
        public static new RunbookProcessStepDeployPackageActionWindowsServiceArgs Empty => new RunbookProcessStepDeployPackageActionWindowsServiceArgs();
    }
}
