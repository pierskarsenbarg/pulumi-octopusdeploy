// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace PiersKarsenbarg.Octopusdeploy.Outputs
{

    [OutputType]
    public sealed class ScriptModuleScript
    {
        /// <summary>
        /// The body of this script module.
        /// </summary>
        public readonly string Body;
        /// <summary>
        /// The syntax of the script. Valid types are `Bash`, `CSharp`, `FSharp`, `PowerShell`, or `Python`.
        /// </summary>
        public readonly string Syntax;

        [OutputConstructor]
        private ScriptModuleScript(
            string body,

            string syntax)
        {
            Body = body;
            Syntax = syntax;
        }
    }
}
